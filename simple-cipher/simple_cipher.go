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

func sanitizeInput(input string) string {
	return reg.ReplaceAllString(strings.ToLower(input), "")
}

type shift struct {
	distance int
}
type vigenere struct {
	key string
}

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
	input = sanitizeInput(input)

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
	if reg.MatchString(key) {
		return nil
	}

	keyLen := len(key)

	if keyLen == 0 || keyLen == 1 {
		return nil
	}

	firstChar := key[0]
	for _, l := range key {
		if l != rune(firstChar) {
			return &vigenere{key: key}
		}
	}

	return nil
}

func (v vigenere) Encode(input string) string {
	input = sanitizeInput(input)

	var w bytes.Buffer
	shiftLen := len(v.key)
	shiftIndex := 0

	for _, letter := range input {
		if shiftLen <= shiftIndex {
			shiftIndex = 0
		}

		var t rune
		shift := letter + rune(v.key[shiftIndex]) - a

		if shift > z {
			t = a + shift - z - 1
		} else if shift < a {
			t = z + shift - a + 1
		} else {
			t = shift
		}
		w.WriteRune(t)
		shiftIndex++
	}

	return w.String()
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}
