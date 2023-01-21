package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outPutDigits []int, err error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	digitsLen := len(inputDigits)
	num := 0

	for i, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}

		gin := math.Pow(float64(inputBase), float64(digitsLen-1-i))

		// println(digitsLen - 1 - i)
		sum := digit * int(gin)
		// fmt.Printf("%d * (%d^%d) = %d\n", digit, inputBase, digitsLen-1-i, sum)
		num += sum
	}
	println(num)

	// slc := []int{}
	// for num > 0 {
	// 	slc = append(slc, num%10)
	// 	num = num / 10
	// }

	// for _, n := range slc {
	// 	println(n)
	// }

	// for i := digitsLen - 1; i >= 0; i-- {
	// 	num += inputDigits[i] * (inputBase ^ i)
	// }

	return
}
