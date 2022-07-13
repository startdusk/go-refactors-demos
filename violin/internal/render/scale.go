package render

import "strings"

// SetDefaultScaleOptions provide the defaults for rendering scales.
func SetDefaultScaleOptions() ([]ScaleOptions, []ScaleOptions, []ScaleOptions, []ScaleOptions) {

	// Set the default scaleOptions for scales and arpeggios.
	sOptions := []ScaleOptions{
		{"Scalearp", "Scale", false, true, "Scales"},
		{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
	}

	// Set the default PitchOptions for scales and arpeggios.
	pOptions := []ScaleOptions{
		{"Pitch", "Major", false, true, "Major"},
		{"Pitch", "Minor", false, false, "Minor"},
	}

	// Set the default KeyOptions for scales and arpeggios.
	kOptions := []ScaleOptions{
		{"Key", "A", false, true, "A"},
		{"Key", "Bb", false, false, "Bb"},
		{"Key", "B", false, false, "B"},
		{"Key", "C", false, false, "C"},
		{"Key", "C#/Db", false, false, "C#/Db"},
		{"Key", "D", false, false, "D"},
		{"Key", "Eb", false, false, "Eb"},
		{"Key", "E", false, false, "E"},
		{"Key", "F", false, false, "F"},
		{"Key", "F#/Gb", false, false, "F#/Gb"},
		{"Key", "G", false, false, "G"},
		{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	// Set the default OctaveOptions for scales and arpeggios.
	oOptions := []ScaleOptions{
		{"Octave", "1", false, true, "1 Octave"},
		{"Octave", "2", false, false, "2 Octave"},
	}
	return sOptions, pOptions, kOptions, oOptions
}

// SetKeyOptions sets the key options based on specified key.
func SetKeyOptions(key string) (kOptions []ScaleOptions) {
	switch key {
	case "A":
		kOptions = []ScaleOptions{
			{"Key", "A", false, true, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Bb":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, true, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "B":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, true, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, true, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C#/Db":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, true, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "D":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, true, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Eb":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, true, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "E":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, true, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, true, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F#/Gb":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, true, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "G":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, true, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}

	case "G#/Ab":
		kOptions = []ScaleOptions{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, true, "G#/Ab"},
		}
	}

	return kOptions
}

// ChangeSharpToS WE DON'T KNOW WHY YET.
func ChangeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}
	return path
}
