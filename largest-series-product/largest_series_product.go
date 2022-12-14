package lsproduct

import (
	"errors"
)

var (
	errLongSpan      = errors.New("span must be smaller than string length")
	errInvalidDigits = errors.New("digits input must only contain digits")
	errNegativeSpan  = errors.New("span must not be negative")
)

func validateInput(digits string, span int) (bool, int64, error) {
	if span > len(digits) {
		return true, 0, errLongSpan
	}
	if span == 0 {
		return true, 1, nil
	}
	if span < 0 {
		return true, 0, errNegativeSpan
	}

	return false, 0, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if invalid, num, err := validateInput(digits, span); invalid {
		return num, err
	}

	zeroProd := true
	var prod int64 = 0

	for i := 0; i < len(digits)-span+1; i++ {
		var newProd int64 = 1
		for _, x := range digits[i : i+span] {
			if x < '0' || x > '9' {
				return 0, errInvalidDigits
			}

			newProd *= int64(x) - '0'
		}

		if newProd > 0 {
			zeroProd = false
		}

		if newProd > prod {
			prod = newProd
		}
	}
	if zeroProd {
		return 0, nil
	}

	return prod, nil
}
