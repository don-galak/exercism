package cryptosquare

import (
	"bytes"
	"math"
	"strings"
)

func Encode(pt string) string {
	var w bytes.Buffer

	for _, r := range strings.ToLower(pt) {
		switch {
		case r >= '0' && r <= '9':
			w.WriteRune(r)
		case r >= 'a' && r <= 'z':
			w.WriteRune(r)
		}
	}
	sqRoot := math.Sqrt(float64(w.Len()))
	lol := []string{}
	var c int
	var r int
	if floored := math.Floor(sqRoot); sqRoot-floored == 0 {
		c = int(sqRoot)
		r = c

	} else {
		c = int(floored) + 1
		r = int(floored)
		println("C: ", c, "R: ", r)
	}
	for w.Len() < c*r {
		w.WriteRune(' ')
	}

	for i := 0; i < w.Len(); i += r {
		lol = append(lol, w.String()[i:i+c-1])
	}
	println(strings.Join(lol, ""), w.String(), pt)
	// for _, k := range lol {
	// 	println(k)
	// }
	// for i := 0; i < c; i++ {
	// 	for j := 0; j < r; j++ {
	// 		lol[i][j] += w.String()[i+j]
	// 	}
	// }

	// println(w.String(), sqRoot)

	return ""
}
