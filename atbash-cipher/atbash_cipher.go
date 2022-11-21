package atbash

import (
	"strconv"
	"strings"
)

const (
	a = 97
	z = 122
	Z = 90
	A = 65
)

func Atbash(s string) string {
	var w strings.Builder
	s = strings.ReplaceAll(s, " .,", "")
	println(s)
	for i, r := range s {
		if _, err := strconv.Atoi(string(r)); err == nil {
			w.WriteRune(r)
		} else if r >= a && r <= z {
			w.WriteRune(z - (r - a))
		} else if r >= A && r <= Z {
			w.WriteRune(Z - (r - A) + 32)
		}
		if i+1%5 == 0 {
			w.WriteString(" ")
		}
	}
	return w.String()
}
