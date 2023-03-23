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

	toAdd := false

	for _, c := range s {
		if toAdd && c != ' ' {
			b.WriteRune(c)
			toAdd = false
		} else if c == ' ' {
			toAdd = true
		}
	}

	return strings.ToUpper(b.String())
}
