package wordsearch

import "fmt"

var steps = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	rows, cols := len(puzzle), len(puzzle[0])
	result := map[string][2][2]int{}
	for _, w := range words {
		res, err := func() ([2][2]int, error) {
			for x, p := range puzzle {
				for y := range p {
					for _, step := range steps {
						outOfBounds := false
						i, j := x, y
						for _, l := range w {
							if isOutOfBounds(i, j, rows, cols) || puzzle[i][j] != byte(l) {
								outOfBounds = true
								break
							}
							i += step[0]
							j += step[1]
						}
						if !outOfBounds {
							return [2][2]int{{y, x}, {j - step[1], i - step[0]}}, nil
						}
					}
				}
			}
			return [2][2]int{}, fmt.Errorf("word %s not found", w)
		}()

		if err != nil {
			return result, err
		}
		result[w] = res
	}
	return result, nil
}

func isOutOfBounds(i, j, rows, cols int) bool {
	if i >= rows || j >= cols || i < 0 || j < 0 {
		return true
	}
	return false
}
