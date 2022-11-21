package atbash

import (
	"strings"
)

func Atbash(s string) string {
	var w strings.Builder
	for _, r := range strings.ToLower(s) {
		if (w.Len()+1)%6 == 0 {
			w.WriteRune(' ')
		}
		switch {
		case r >= '0' && r <= '9':
			w.WriteRune(r)
		case r >= 'a' && r <= 'z':
			w.WriteRune('z' - (r - 'a'))
		}
	}
	return strings.TrimRight(w.String(), " ")
}
