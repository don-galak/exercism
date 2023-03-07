package resistorcolorduo

var colorSlice = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func findColorIndex(color string) int {
	for i, c := range colorSlice {
		if c == color {
			return i
		}
	}
	return -1
}

func Value(colors []string) int {
	return findColorIndex(colors[0])*10 + findColorIndex(colors[1])
}
