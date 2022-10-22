package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	data [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	var data [][]int

	for i, row := range rows {
		r := strings.Fields(row)
		if i > 0 && len(r) != len(data[i-1]) {
			return nil, errors.New("matrix rows have inconsistent length")
		}

		data = append(data, []int{})
		for _, char := range r {
			value, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			data[i] = append(data[i], value)
		}
	}

	return &Matrix{data: data}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() (cols [][]int) {
	cols = make([][]int, len(m.data[0]))
	for j := 0; j < len(m.data); j++ {
		for i := 0; i < len(cols); i++ {
			cols[i] = append(cols[i], m.data[j][i])
		}
	}

	return
}

func (m *Matrix) Rows() (rows [][]int) {
	rows = make([][]int, len(m.data))
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(m.data[i]); j++ {
			rows[i] = append(rows[i], m.data[i][j])
		}
	}

	return
}

func (m *Matrix) Set(row, col, val int) bool {
	rowLen := len(m.data)
	colLen := len(m.data[0])

	if row < 0 || col < 0 || row >= rowLen || col >= colLen {
		return false
	}

	m.data[row][col] = val
	return true
}
