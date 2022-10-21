package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	data [][]int
}

func New(s string) (*Matrix, error) {
	splitted := strings.Split(s, "\n")
	var data [][]int

	for i, r := range splitted {
		rowInputs := strings.Split(r, " ")
		data = append(data, []int{})
		for _, char := range rowInputs {
			if value, err := strconv.Atoi(strings.TrimLeft(string(char), " ")); err == nil {
				data[i] = append(data[i], value)
			}
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
	panic("Please implement the Set function")
}
