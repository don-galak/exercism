package darts

func Score(x, y float64) int {
	hypotenuse := (x * x) + (y * y)
	switch {
	case hypotenuse <= 1:
		return 10
	case hypotenuse <= 25:
		return 5
	case hypotenuse <= 100:
		return 1
	default:
		return 0
	}
}
