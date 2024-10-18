package cryptosquare

import (
	"bytes"
	"math"
	"strings"
)

func Encode(pt string) string {
	if pt == "" {
		return ""
	}
	var w bytes.Buffer

	for _, r := range strings.ToLower(pt) {
		switch {
		case r >= '0' && r <= '9':
			w.WriteRune(r)
		case r >= 'a' && r <= 'z':
			w.WriteRune(r)
		}
	}

	l := w.Len()
	c := int(math.Ceil(math.Sqrt(float64(l))))
	r := int(math.Ceil(float64(l) / float64(c)))
	byt := w.Bytes()

	matrix := make([][]rune, r)
	for i := 0; i < r; i++ {
		matrix[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			if i*c+j < l {
				matrix[i][j] = rune(byt[i*c+j])
			} else {
				matrix[i][j] = ' '
			}
		}
	}

	w.Reset()
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			w.WriteRune(rune(matrix[i][j]))
		}
		if j+1 < c {
			w.WriteRune(' ')
		}
	}

	return w.String()
}
