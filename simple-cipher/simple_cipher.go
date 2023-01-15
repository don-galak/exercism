package cipher

import "bytes"

const (
	a = 'a'
	z = 'z'
)

type shift struct {
	distance int
}
type vigenere struct{}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	return &shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var w bytes.Buffer

	for _, letter := range input {
		var t rune
		if sum := letter + rune(c.distance); sum > z {
			t = a + sum - z - 1
		} else {
			t = sum
		}
		w.WriteRune(t)
	}
	return w.String()
}

func (c shift) Decode(input string) string {
	// c.distance = -c.distance
	// return c.Encode(input)
	var w bytes.Buffer

	for _, letter := range input {
		var t rune
		if diff := letter - rune(c.distance); diff < a {
			t = z + diff - a + 1
		} else {
			t = diff
		}
		w.WriteRune(t)
	}
	return w.String()
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
