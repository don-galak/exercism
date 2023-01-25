package perfect

import (
	"errors"
)

type Classification string

const (
	ClassificationAbundant  = "ClassificationAbundant"
	ClassificationDeficient = "ClassificationDeficient"
	ClassificationPerfect   = "ClassificationPerfect"
)

var ErrOnlyPositive = errors.New("input is not a positive integer")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	var sum int64 = 0

	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}
