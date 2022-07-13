package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"violin/cmd/violin/internal/handlers"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	// =======================================================================================
	// Logging

	log := log.New(os.Stdout, "VIOLIN", log.Lshortfile|log.Ldate|log.Lmicroseconds)

	// =======================================================================================
	// Configuration

	var cfg struct {
		Web struct {
			APIHost         string        `conf:"default:0.0.0.0:8080"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}
	if err := conf.Parse(os.Args[1:], "VIOLIN", &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			usage, err := conf.Usage("VIOLIN", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config error")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	log.Println("main : Started : Application initializing")
	defer log.Println("main : Completed")

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main : Config :\n%v\n", out)

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      handlers.NewMux(log),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	// Mark a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// =======================================================================================
	// Configuration

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")
	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", cfg.Web.ShutdownTimeout, err)
			err = api.Close()
		}

		// log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.Wrap(err, "integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}
	return nil
}
