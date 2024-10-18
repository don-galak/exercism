package rotationalcipher

import (
	"strings"
)

const (
	a = 97
	z = 122
	Z = 90
	A = 65
)

func RotationalCipher(plain string, shiftKey int) string {
	var w strings.Builder
	if shiftKey == 26 {
		shiftKey = 0
	}

	for _, r := range plain {
		shifted := r + rune(shiftKey)
		if r >= a && r <= z {
			if shifted > z {
				shifted = a + (shifted - z - 1)
			}
		} else if r >= A && r <= Z {
			if shifted > Z {
				shifted = A + (shifted - Z - 1)
			}
		} else {
			shifted = r
		}
		w.WriteRune(shifted)
	}

	return w.String()
}
