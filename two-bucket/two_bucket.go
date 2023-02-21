package twobucket

import (
	"errors"
	"fmt"
)

const (
	one = "one"
	two = "two"
)

var errInvalidInput = errors.New("invalid")
var invalid = func() (string, int, int, error) { return "", 0, 0, errInvalidInput }

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne == 0 || sizeBucketTwo == 0 || goalAmount == 0 {
		return invalid()
	}

	moves := 0
	bucketOne := 0
	bucketTwo := 0

	if startBucket == one {
		//fill bucket one
		bucketOne = sizeBucketOne
		moves++

		//pour bucket one in bucket two
		for bucketOne > 0 {
			bucketOne--
			bucketTwo++
		}
		moves++

		// fill  bucket one
		bucketOne = sizeBucketOne
		moves++

		//pour bucket one in bucket two
		for bucketTwo < sizeBucketTwo {
			bucketOne--
			bucketTwo++
		}
		moves++

		return one, moves, bucketTwo, nil
	}

	if startBucket == two {
		//fill bucket two
		bucketTwo = sizeBucketTwo
		moves++
		printer(moves, bucketOne, bucketTwo)

		//pour bucket two in bucket one
		for bucketOne < sizeBucketOne {
			bucketTwo--
			bucketOne++
		}
		moves++
		printer(moves, bucketOne, bucketTwo)

		// empty bucket one
		bucketOne = 0
		moves++
		printer(moves, bucketOne, bucketTwo)

		//pour bucket two in bucket one
		for bucketTwo > 0 {
			bucketOne++
			bucketTwo--
		}
		moves++
		printer(moves, bucketOne, bucketTwo)

		//fill bucket two
		bucketTwo = sizeBucketTwo
		moves++
		printer(moves, bucketOne, bucketTwo)

		//pour bucket two in bucket one
		for bucketOne < sizeBucketOne {
			bucketOne++
			bucketTwo--
		}
		moves++
		printer(moves, bucketOne, bucketTwo)

		// empty bucket one
		bucketOne = 0
		moves++
		printer(moves, bucketOne, bucketTwo)

		//pour bucket two in bucket one
		for bucketOne < sizeBucketOne {
			bucketOne++
			bucketTwo--
		}
		moves++
		printer(moves, bucketOne, bucketTwo)

		// return two, moves, bucketOne, nil
	}

	return invalid()
}

func printer(step, bucketOne, bucketTwo int) {
	fmt.Printf("step: %d, bucketOne: %d, bucketTwo: %d\n\n", step, bucketOne, bucketTwo)
}
