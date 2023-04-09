package wordsearch

import (
	"strings"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, w := range words {
		result[w] = [2][2]int{{-1, -1}, {-1, -1}}
	}

	for i, p := range puzzle {
		for _, w := range words {
			first := strings.Index(p, w)
			if first == -1 {
				continue
			}
			last := first + len(w) - 1
			result[w] = [2][2]int{{first, i}, {last, i}}
		}
	}

	return result, err
}

// make [][]slice of input puzzle
// range over and find first letter of first word in slice
// find at which direction lies the next letter
// right, left, top, bottom, top-right, top-left, bottom-right, bottom-left
// increment/decrement i,j as steps accordingly
// repeat until word is complete
// repeat steps for next word