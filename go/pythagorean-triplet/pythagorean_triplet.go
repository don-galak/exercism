package pythagorean

import "math"

type Triplet [3]int

func (t Triplet) qualifiesSum(p int) bool {
	return t[0]+t[1]+t[2] == p
}

func getPythagoreanTriplet(x, y, z int) (bool, Triplet) {
	if x == y+z || y == x+z || z == y+x {
		return true, Triplet{int((math.Sqrt(float64(x)))), int((math.Sqrt(float64(y)))), int((math.Sqrt(float64(z))))}
	}
	return false, Triplet{}
}

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	intArr := []int{}
	for i := min; i < max+1; i++ {
		intArr = append(intArr, i*i)
	}
	triplets := []Triplet{}

	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			for k := j + 1; k < len(intArr); k++ {
				if f, triplet := getPythagoreanTriplet(intArr[i], intArr[j], intArr[k]); f {
					triplets = append(triplets, triplet)
				}
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (triplets []Triplet) {
	for _, triplet := range Range(0, p/2) {
		if triplet.qualifiesSum(p) {
			triplets = append(triplets, triplet)
		}
	}
	return
}
