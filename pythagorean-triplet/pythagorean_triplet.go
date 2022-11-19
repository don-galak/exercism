package pythagorean

type Triplet [3]int

func (t Triplet) qualifiesSum(p int) bool {
	if t[0]+t[1]+t[2] == p {
		return true
	}
	return false
}

func isPythagoreanTriplet(x, y, z int) bool {
	if x == y+z || y == x+z || z == y+x {
		return true
	}
	return false
}

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	arr := []int{}
	for i := min; i < max+1; i++ {
		arr = append(arr, i)
	}
	triplets := []Triplet{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				x := arr[i] * arr[i]
				y := arr[j] * arr[j]
				z := arr[k] * arr[k]
				if isPythagoreanTriplet(x, y, z) {
					triplets = append(triplets, Triplet{arr[i], arr[j], arr[k]})
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
