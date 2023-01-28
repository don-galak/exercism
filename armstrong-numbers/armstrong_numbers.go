package armstrong

import "math"

func IsNumber(n int) bool {
	if n < 10 {
		return true
	}

	numOfDigits := getNumberOfDigits(n)
	sum := getSum(n, numOfDigits)

	return n == sum
}

func getNumberOfDigits(n int) (out int) {
	for n > 0 {
		out++
		if n < 10 {
			break
		}
		n /= 10
	}
	return out
}

func getSum(n, numOfDigits int) (sum int) {
	for n > 0 {
		sum += int(math.Pow(float64(n%10), float64(numOfDigits)))
		if n < 10 {
			break
		}
		n /= 10
	}
	return
}
