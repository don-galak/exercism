package diamond

import (
	"errors"
	"strings"
)

var errOutOfRange = errors.New("out of range")

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errOutOfRange
	}

	start := 'A'
	sliceLen := int(char) - 64
	slice := make([]string, sliceLen)
	sliceLen--

	slice[0] = strings.Repeat(" ", sliceLen)
	slice[0] += string(start)
	slice[0] += strings.Repeat(" ", sliceLen)

	between := 1
	sliceLen--
	start++
	for i := 1; i < len(slice); i++ {
		slice[i] = strings.Repeat(" ", sliceLen)
		slice[i] += string(start)
		slice[i] += strings.Repeat(" ", between)
		slice[i] += string(start)
		slice[i] += strings.Repeat(" ", sliceLen)
		between += 2
		sliceLen--
		start++
	}

	for i := len(slice) - 2; i > -1; i-- {
		slice = append(slice, slice[i])
	}

	return strings.Join(slice, "\n"), nil
}
