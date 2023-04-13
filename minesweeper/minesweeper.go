package minesweeper

import (
	"fmt"
	"strconv"
	"strings"
)

var directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Annotate(board []string) []string {
	rows := len(board)
	if rows == 0 || len(board[0]) == 1 {
		return board
	}

	SECOND := make([][]string, len(board))
	for i := range SECOND {
		SECOND[i] = make([]string, len(board[0]))
	}

	for i, row := range board {
		for j, c := range row {
			SECOND[i][j] = string(c)
		}
	}
	cols := len(board[0])

	for i, row := range SECOND {
		for j := range row {
			for _, d := range directions {
				x, y := i, j
				x += d[0]
				y += d[1]
				if x >= rows || y >= cols || x < 0 || y < 0 {
					continue
				}

				if SECOND[x][y] == "*" {
					if SECOND[i][j] == " " {
						SECOND[i][j] = "1"
						continue
					}

					if n, err := strconv.Atoi(SECOND[i][j]); err == nil {
						SECOND[i][j] = fmt.Sprintf("%d", n+1)
					}
				}
			}
		}
	}

	for i, row := range SECOND {
		board[i] = strings.Join(row, "")
	}

	return board
}
