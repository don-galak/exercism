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

func (m *Matrix) isSaddlePoint(i, j int) bool {
	for c := range (*m)[0] {
		if (*m)[i][c] > (*m)[i][j] {
			return false
		}
	}
	for r := range *m {
		if (*m)[r][j] < (*m)[i][j] {
			return false
		}
	}
	return true
}
func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for i, row := range *m {
		for j := range row {
			if m.isSaddlePoint(i, j) {
				pairs = append(pairs, Pair{i + 1, j + 1})
			}
		}
	}
	return pairs
}
