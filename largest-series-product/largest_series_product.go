package lsproduct

import (
	"errors"
	"regexp"
)

var (
	errLongSpan      = errors.New("span must be smaller than string length")
	errInvalidDigits = errors.New("digits input must only contain digits")
	errNegativeSpan  = errors.New("span must not be negative")
	re               = regexp.MustCompile(`^\d+$`)
)

func sanitizeInput(digits string, span int) (bool, int64, error) {
	if span > len(digits) {
		return true, 0, errLongSpan
	}
	if span == 0 {
		return true, 1, nil
	}
	if !re.MatchString(digits) {
		return true, 0, errInvalidDigits
	}
	if span < 0 {
		return true, 0, errNegativeSpan
	}

	return false, 0, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if invalid, num, err := sanitizeInput(digits, span); invalid {
		return num, err
	}

	zeroProd := true
	var prod int64 = 1

	for i := 0; i < len(digits)-span+1; i++ {
		var newProd int64 = 1
		for _, x := range digits[i : i+span] {
			newProd *= int64(x) - '0'
		}

		if newProd > 0 {
			zeroProd = false
		}

		if prod < newProd {
			prod = newProd
		}
	}
	if zeroProd {
		return 0, nil
	}

	return prod, nil
}
