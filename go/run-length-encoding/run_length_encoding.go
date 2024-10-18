package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	var w strings.Builder
	for len(s) > 0 {
		letter := s[0]
		slen := len(s)
		s = strings.TrimLeft(s, string(letter))
		if n := slen - len(s); n > 1 {
			w.WriteString(strconv.Itoa(n))
		}
		w.WriteString(string(letter))
	}
	return w.String()
}

func RunLengthDecode(s string) string {
	var w strings.Builder
	for len(s) > 0 {
		i := strings.IndexFunc(s, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(s[:i])
		}
		w.WriteString(strings.Repeat(string(s[i]), n))
		s = s[i+1:]
	}
	return w.String()
}
