package minesweeper

import (
	"fmt"
	"strconv"
	"strings"
)

var directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Annotate(board []string) []string {
	rows := len(board)
	if rows == 0 {
		return board
	}

	temp := make([][]string, len(board))
	for i := range temp {
		temp[i] = make([]string, len(board[0]))
	}

	for i, row := range board {
		for j, c := range row {
			temp[i][j] = string(c)
		}
	}
	cols := len(board[0])

	for i, row := range temp {
		for j := range row {
			for _, d := range directions {
				x, y := i, j
				x += d[0]
				y += d[1]
				if x >= rows || y >= cols || x < 0 || y < 0 {
					continue
				}

				if temp[x][y] == "*" {
					if temp[i][j] == " " {
						temp[i][j] = "1"
						continue
					}
					if n, err := strconv.Atoi(temp[i][j]); err == nil {
						temp[i][j] = fmt.Sprintf("%d", n+1)
					}
				}
			}
		}
	}

	annotated := make([]string, len(board))
	for i, row := range temp {
		annotated[i] = strings.Join(row, "")
	}

	return annotated
}
