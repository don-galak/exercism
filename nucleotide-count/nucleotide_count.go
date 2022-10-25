package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range d {
		if _, exists := h[n]; !exists {
			return h, errors.New("invalid input")
		}
		h[n]++
	}
	return h, nil
}
