package acronym

import (
	"bytes"
	"strings"
)

func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")

	var b bytes.Buffer
	b.WriteByte(s[0])

	write := false

	for _, c := range s {
		if write && c != ' ' {
			b.WriteRune(c)
			write = false
		} else if c == ' ' {
			write = true
		}
	}

	return strings.ToUpper(b.String())
}
