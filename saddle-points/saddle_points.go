package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Pair struct{ i, j int }
type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	var matrix Matrix

	for i, row := range rows {
		r := strings.Fields(row)
		if i > 0 && len(r) != len(matrix[i-1]) {
			return nil, errors.New("matrix rows have inconsistent length")
		}

		matrix = append(matrix, []int{})
		for _, char := range r {
			value, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			matrix[i] = append(matrix[i], value)
		}
	}

	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for i, r := range *m {
		for j, el := range r {
			saddleInRow := true
			for x, rowItem := range (*m)[i] {
				if x != j && rowItem > el {
					saddleInRow = false
					break
				}
			}
			if !saddleInRow {
				continue
			}
			saddleInCol := true
			for y, col := range *m {
				if y != i && col[j] < el {
					saddleInCol = false
					break
				}
			}
			if !saddleInCol {
				continue
			}

			pairs = append(pairs, Pair{i + 1, j + 1})
		}
	}
	return pairs
}
