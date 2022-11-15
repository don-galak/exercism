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

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return 0, errLongSpan
	}
	if span == 0 {
		return 1, nil
	}
	if !re.Match([]byte(digits)) {
		return 0, errInvalidDigits
	}
	if span < 0 {
		return 0, errNegativeSpan
	}

	zeroProd := true
	var prod int64 = 1
	counter := len(digits) - span + 1

	for i := 0; i < counter; i++ {
		var newProd int64 = 1
		for _, x := range digits[i : i+span] {
			newProd *= int64(x - '0')
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
