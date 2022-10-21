package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows [][]int
	cols [][]int
}

func New(s string) (*Matrix, error) {
	splitted := strings.Split(s, "\n")
	var cols [][]int
	var rows [][]int

	for i, r := range splitted {
		rowInputs := strings.Split(r, " ")
		rows = append(rows, []int{})
		for _, rowInput := range rowInputs {
			value, err := strconv.Atoi(rowInput)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			rows[i] = append(rows[i], value)
		}
	}

	for i := range rows {
		for j := range rows[i] {
			cols = append(cols, []int{})
			cols[j] = append(cols[j], rows[i][j])
		}
	}

	return &Matrix{rows: rows, cols: cols}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	return append(make([][]int, 0, len(m.cols)), m.cols...)
}

func (m *Matrix) Rows() [][]int {
	return append(make([][]int, 0, len(m.rows)), m.rows...)
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
