package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
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

// Cols and Rows must return the results without affecting the matrix.
func (matrix *Matrix) Cols() (cols [][]int) {
	cols = make([][]int, len((*matrix)[0]))
	for j := 0; j < len(*matrix); j++ {
		for i := 0; i < len(cols); i++ {
			cols[i] = append(cols[i], (*matrix)[j][i])
		}
	}

	return
}

func (matrix *Matrix) Rows() (rows [][]int) {
	rows = make([][]int, len(*matrix))
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len((*matrix)[i]); j++ {
			rows[i] = append(rows[i], (*matrix)[i][j])
		}
	}

	return
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(*m) || col >= len((*m)[0]) {
		return false
	}

	(*m)[row][col] = val
	return true
}
