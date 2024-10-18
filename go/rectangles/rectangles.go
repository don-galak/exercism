package rectangles

// Count the rectangles in an ASCII diagram.
func Count(diagram []string) int {
	height := len(diagram)
	if height == 0 {
		return 0
	}
	width := len(diagram[0])
	count := 0
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			if diagram[y][x] == '+' {
				for xx := x + 1; xx < width && (diagram[y][xx] == '-' || diagram[y][xx] == '+'); xx++ {
					if diagram[y][xx] == '+' {
						for yy := y + 1; yy < height && (diagram[yy][xx] == '|' || diagram[yy][xx] == '+') && (diagram[yy][x] == '|' || diagram[yy][x] == '+'); yy++ {
							if diagram[yy][xx] == '+' && diagram[yy][x] == '+' {
								xxx := x + 1
								for ; xxx < xx && (diagram[yy][xxx] == '-' || diagram[yy][xxx] == '+'); xxx++ {
								}
								if xxx == xx {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	return count
}
