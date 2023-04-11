package wordsearch

import (
	"errors"
	"strings"
)

type direction struct {
	i int
	j int
}

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
				// result[w] = [2][2]int{{firstLetterIndex, x}, {lastLetterIndex, x}}
				// println(x, firstLetterIndex)
				// dir := direction{0, 0}
				availableCells := getAvailableCellsBasedOnCollision(x, firstLetterIndex, rows, cols)

				for _, cell := range availableCells {
					i := cell[0]
					j := cell[1]
					println(i, j)
					if string(w[1]) == string(puzzle[i][j]) {
						println("TO BRIKAME")
					}
				}
			}
		}
	}
	if !foundWord {
		err = errors.New("")

	}

	return result, err
}

func getAvailableCellsBasedOnCollision(x, y, rows, cols int) [][2]int {
	collisionRight := y == cols
	collisionLeft := y == 0
	collisionTop := x == 0
	collisionBottom := x+1 >= rows

	if collisionBottom && collisionTop && !collisionLeft && !collisionRight {
		return [][2]int{{0, y + 1}, {0, y - 1}}
	}
	if collisionBottom && collisionTop && collisionLeft && !collisionRight {
		return [][2]int{{0, y + 1}}
	}
	if collisionBottom && collisionTop && !collisionLeft && collisionRight {
		return [][2]int{{0, y - 1}}
	}
	if collisionRight && collisionBottom && !collisionTop {
		return [][2]int{{x - 1, y}, {x, y - 1}, {x - 1, y - 1}}
	}
	if collisionLeft && collisionBottom && !collisionTop {
		return [][2]int{{x - 1, y}, {x, y + 1}, {x - 1, y + 1}}
	}
	if collisionTop && collisionLeft && !collisionBottom {
		return [][2]int{{x + 1, y}, {x, y + 1}, {x + 1, y + 1}}
	}
	if collisionTop && collisionRight && !collisionBottom {
		return [][2]int{{x + 1, y}, {x, y - 1}, {x + 1, y - 1}}
	}
	if collisionLeft && !collisionBottom && !collisionTop {
		return [][2]int{{x - 1, y}, {x + 1, y + 1}, {x, y + 1}, {x + 1, y + 1}, {x + 1, y}}
	}
	if collisionRight && !collisionBottom && !collisionTop {
		return [][2]int{{x - 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}}
	}
	if collisionTop && !collisionRight && collisionLeft {
		return [][2]int{{x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}, {x, y + 1}}
	}
	if collisionBottom && !collisionRight && collisionLeft {
		return [][2]int{{x, y - 1}, {x - 1, y - 1}, {x - 1, y}, {x - 1, y - 1}, {x, y + 1}}
	}
	return [][2]int{{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}, {x, y + 1}}
}

// make [][]slice of input puzzle
// range over and find first letter of first word in slice
// find at which direction lies the next letter
// right, left, top, bottom, top-right, top-left, bottom-right, bottom-left
// increment/decrement i,j as steps accordingly
// repeat until word is complete
// repeat steps for next word
