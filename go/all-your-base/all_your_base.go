package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (out []int, err error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	sum := 0

	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}

		sum = sum*inputBase + digit
	}

	if sum == 0 {
		return []int{0}, nil
	}

	for ; sum > 0; sum /= outputBase {
		out = append(out, sum%outputBase)
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return
}
