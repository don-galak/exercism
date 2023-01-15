package cipher

import (
	"bytes"
	"regexp"
	"strings"
)

const (
	a = 'a'
	z = 'z'
)

var reg = regexp.MustCompile(`[^a-z]+`)

type shift struct {
	distance int
}
type vigenere struct{}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance > 25 || distance < -25 || distance == 0 {
		return nil
	}

	return &shift{distance: distance}
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	input = reg.ReplaceAllString(input, "")

	var w bytes.Buffer

	for _, letter := range input {
		var t rune
		sum := letter + rune(c.distance)

		if sum > z {
			t = a + sum - z - 1
		} else if sum < a {
			t = z + sum - a + 1
		} else {
			t = sum
		}
		w.WriteRune(t)
	}
	return w.String()
}

func (c shift) Decode(input string) string {
	c.distance = -c.distance
	return c.Encode(input)
}

func NewVigenere(key string) Cipher {
	panic("Please implement the NewVigenere function")
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}
