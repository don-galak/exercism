package armstrong

import "math"

func IsNumber(n int) bool {
	if n < 10 {
		return true
	}
	return n == getSum(n)
}

func getSum(n int) (sum int) {
	digits := getNumberOfDigits(n)
	for n > 0 {
		sum += int(math.Pow(float64(n%10), float64(digits)))
		n /= 10
	}
	return
}

func getNumberOfDigits(n int) int {
	var i int
	for i = 0; n > 0; i++ {
		n /= 10
	}
	return i
}
