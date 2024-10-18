package resistorcolortrio

import (
	"fmt"
	"math"
)

var colorSlice = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
var magnitudes = []string{"", "kilo", "mega", "giga"}

func findColorIndex(color string) int {
	for i, c := range colorSlice {
		if c == color {
			return i
		}
	}
	return -1
}

func Label(colors []string) string {
	c1 := findColorIndex(colors[0])
	c2 := findColorIndex(colors[1])
	zeroes := math.Pow10(findColorIndex(colors[2]))
	resistance := (c1*10 + c2) * int(zeroes)
	mag := 0

	for ; resistance >= 1000; resistance /= 1000 {
		mag++
	}
	return fmt.Sprintf("%d %sohms", resistance, magnitudes[mag])
}
