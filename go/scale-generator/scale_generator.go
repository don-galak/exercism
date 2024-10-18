package scale

import "strings"

var steps = map[rune]int{'m': 1, 'M': 2, 'A': 3}

func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmm"
	}
	interval += "m"
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		notes = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}
	var index int
	for i := range notes {
		if strings.EqualFold(notes[i], tonic) {
			index = i
			break
		}
	}
	scale := make([]string, len(interval))
	for i, v := range interval {
		scale[i] = notes[(index)%len(notes)]
		index += steps[v]
	}
	return scale
}
