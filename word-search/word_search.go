package wordsearch

import (
	"errors"
	"strings"
)

type direction struct {
	i int
	j int
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, w := range words {
		result[w] = [2][2]int{{-1, -1}, {-1, -1}}
	}

	// rows := len(puzzle)
	// cols := len(puzzle[0])

	for _, w := range words {
		for x, p := range puzzle {
			firstLetterIndex := strings.Index(p, string(w[0]))
			if firstLetterIndex > -1 {
				lastLetterIndex := strings.LastIndex(p, string(w[len(w)-1]))
				result[w] = [2][2]int{{firstLetterIndex, x}, {lastLetterIndex, x}}
				// println(x, firstLetterIndex)
				// dir := direction{-1, -1}
				// if x+1 >= rows || x-1 < 0 {
				// 	dir.i = 0
				// }
				// if firstLetterIndex+1 >= cols || firstLetterIndex-1 < 0 {
				// 	dir.j = 0
				// }

				break
			}
			if firstLetterIndex < 0 {
				err = errors.New("")
			}
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
