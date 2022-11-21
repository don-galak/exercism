package atbash

import (
	"regexp"
	"strings"
)

const (
	zero = 48
	nine = 57
	a    = 97
	z    = 122
	Z    = 90
	A    = 65
)

var re = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func Atbash(s string) string {
	var w strings.Builder
	s = re.ReplaceAllString(s, "")
	for i, r := range s {
		if i != 0 && i%5 == 0 {
			w.WriteRune(' ')
		}
		switch {
		case r >= zero && r <= nine:
			w.WriteRune(r)
		case r >= a && r <= z:
			w.WriteRune(z - (r - a))
		case r >= A && r <= Z:
			w.WriteRune(Z - (r - A) + 32)
		}
	}
	return w.String()
}
