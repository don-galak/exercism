package wordsearch

import (
	"errors"
	"strings"
)

const cols = 9

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	foundWord := false
	for _, w := range words {
		result[w] = [2][2]int{{-1, -1}, {-1, -1}}
	}

	rows := len(puzzle)

	for _, w := range words {
		for x, p := range puzzle {
			firstLetterIndex := strings.Index(p, string(w[0]))
			if firstLetterIndex > -1 {

				steps := getAvailableStepsBasedOnCollision(x, firstLetterIndex, rows, cols)

				for _, step := range steps {
					xStep := step[0]
					yStep := step[1]
					i := x
					j := firstLetterIndex

					possibleMatch := ""

					for k := 0; k < len(w); k++ {
						if isOutOfBounds(i, j, rows, cols) {
							break
						}

						if w[k] == puzzle[i][j] {
							possibleMatch += string(puzzle[i][j])
							if possibleMatch == w {
								result[w] = [2][2]int{{firstLetterIndex, x}, {j, i}}
								println("FOUND IT")
								foundWord = true
								break
							}
							i += xStep
							j += yStep
							continue
						}
						break
					}
				}
			}
		}
	}
	if !foundWord {
		// err = nil
		err = errors.New("")
	}

	return result, err
}

func isOutOfBounds(i, j, rows, cols int) bool {
	if i > rows-1 || j > cols-1 || i < 0 || j < 0 {
		return true
	}
	return false
}

func getAvailableStepsBasedOnCollision(x, y, rows, cols int) [][2]int {
	collisionRight := y == cols
	collisionLeft := y == 0
	collisionTop := x == 0
	collisionBottom := x+1 >= rows

	if collisionBottom && collisionTop && !collisionLeft && !collisionRight {
		return [][2]int{{0, 1}, {0, -1}}
	}
	if collisionBottom && collisionTop && collisionLeft && !collisionRight {
		return [][2]int{{0, 1}}
	}
	if collisionBottom && collisionTop && !collisionLeft && collisionRight {
		return [][2]int{{0, -1}}
	}
	if collisionRight && collisionBottom && !collisionTop {
		return [][2]int{{-1, 0}, {0, -1}, {-1, -1}}
	}
	if collisionLeft && collisionBottom && !collisionTop {
		return [][2]int{{-1, 0}, {0, 1}, {-1, 1}}
	}
	if collisionTop && collisionLeft && !collisionBottom {
		return [][2]int{{1, 0}, {0, 1}, {1, 1}}
	}
	if collisionTop && collisionRight && !collisionBottom {
		return [][2]int{{1, 0}, {0, -1}, {1, -1}}
	}
	if collisionLeft && !collisionBottom && !collisionTop {
		return [][2]int{{-1, 0}, {1, 1}, {0, 1}, {1, 1}, {1, 0}}
	}
	if collisionRight && !collisionBottom && !collisionTop {
		return [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}}
	}
	if collisionTop && !collisionRight && collisionLeft {
		return [][2]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}}
	}
	if collisionBottom && !collisionRight && collisionLeft {
		return [][2]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, -1}, {0, 1}}
	}
	return [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
}

// make [][]slice of input puzzle
// range over and find first letter of first word in slice
// find at which direction lies the next letter
// right, left, top, bottom, t-right, top-left, bottom-right, bottom-left
// increment/decrement i,j as steps accordingly
// repeat until word is complete
// repeat steps for next word
