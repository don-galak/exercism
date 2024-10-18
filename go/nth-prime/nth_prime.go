package prime

import (
	"errors"
)

func isPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i < num>>1; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}

	num := 0
	i := 2
	for n > 0 {
		if isPrime(i) {
			num = i
			n--
		}
		i++
	}

	return num, nil
}
