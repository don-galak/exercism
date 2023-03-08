package resistorcolortrio

import (
	"fmt"
	"math"
)

const (
	ohms, kilo, mega, giga = "ohms", "kilo", "mega", "giga"
)
const (
	KO, MO, GO = 1000, KO * KO, MO * KO
)

var colorSlice = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func findColorIndex(color string) int {
	for i, c := range colorSlice {
		if c == color {
			return i
		}
	}
	return -1
}

func Value(colors []string) string {
	zeroes := findColorIndex(colors[2])
	res := math.Pow(10, float64(zeroes))
	final := (findColorIndex(colors[0])*10 + findColorIndex(colors[1])) * int(res)

	if final == 0 {
		return "0 ohms"
	}

	var suffix string
	switch {
	case final%GO == 0:
		suffix = giga
		final /= GO
	case final%MO == 0:
		suffix = mega
		final /= MO
	case final%KO == 0:
		suffix = kilo
		final /= KO
	default:
		suffix = ""
	}
	suffix += ohms

	return fmt.Sprintf("%d %v", final, suffix)
}

func Label(colors []string) string {
	return Value(colors)
}
