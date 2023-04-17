package transpose

import (
	"bytes"
	"strings"
)

func Transpose(input []string) []string {
	s := make([]string, len(input))

	var b bytes.Buffer
	biggestRowBellowCurrentRow := 0
	biggestRow := 0
	for i, row := range input {
		for x := i + 1; x < len(input); x++ {
			if l := len(input[x]); l > biggestRowBellowCurrentRow {
				biggestRowBellowCurrentRow = l
			}
		}
		b.WriteString(row)
		rowLen := len(row)
		if rowLen > biggestRow {
			biggestRow = rowLen
		}

		if rowLen < biggestRowBellowCurrentRow {
			b.WriteString(strings.Repeat(" ", biggestRowBellowCurrentRow-rowLen))
		}

		s[i] = b.String()
		b.Reset()
		biggestRowBellowCurrentRow = 0
	}
	transposed := make([]string, biggestRow)

	for j := 0; j < len(transposed); j++ {
		for i := 0; i < len(input); i++ {
			if j >= len(s[i]) || s[i][j] == '&' {
				break
			}
			b.WriteByte(s[i][j])
		}
		transposed[j] = b.String()
		b.Reset()
	}

	return transposed
}
