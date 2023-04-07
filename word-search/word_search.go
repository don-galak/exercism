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
				// err = errors.New("")
				continue
			}
			last := first + len(w) - 1
			result[w] = [2][2]int{{first, i}, {last, i}}
		}
	}

	// i := strings.Index("ela mesa", "mesa")
	// j := i + len("mesa") - 1

	// println(i, j)

	return result, err
}
