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

	letter := 'A'
	spacesAround := int(char) - 64
	slice := make([]string, spacesAround)
	spacesAround--

	slice[0] = strings.Repeat(" ", spacesAround)
	slice[0] += string(letter)
	slice[0] += strings.Repeat(" ", spacesAround)

	spacesBetween := 1
	spacesAround--
	letter++
	for i := 1; i < len(slice); i++ {
		slice[i] = strings.Repeat(" ", spacesAround)
		slice[i] += string(letter)
		slice[i] += strings.Repeat(" ", spacesBetween)
		slice[i] += string(letter)
		slice[i] += strings.Repeat(" ", spacesAround)
		spacesBetween += 2
		spacesAround--
		letter++
	}

	for i := len(slice) - 2; i > -1; i-- {
		slice = append(slice, slice[i])
	}

	return strings.Join(slice, "\n"), nil
}
