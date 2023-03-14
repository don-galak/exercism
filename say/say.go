package say

import "bytes"

const (
	ten      = 10
	hundred  = 10 * ten
	thousand = 10 * hundred
	million  = 1000 * thousand
	billion  = 1000 * million
	max      = 1000*billion - 1
)

var m = map[int64]string{0: "zero", 1: "one", 2: "two",
	3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
	8: "eight", 9: "nine", ten: "ten", 11: "eleven", 12: "twelve",
	13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen",
	18: "eighteen", 19: "nineteen",
	20:       "twenty",
	30:       "thirty",
	40:       "fourty",
	50:       "fifty",
	60:       "sixty",
	70:       "seventy",
	80:       "eighty",
	90:       "ninety",
	hundred:  "hundred",
	thousand: "thousand",
	million:  "million",
	billion:  "billion"}

func Say(n int64) (string, bool) {
	if n > max || n < 0 {
		return "", false
	}
	if s, exists := m[n]; exists {
		if n >= hundred {
			return "one " + s, true
		}

		return s, true
	}

	// slice := []int{}

	// for n > 0 {
	// 	println(n, (n/10)*10)
	// 	n /= 10
	// }

	// println((billion+123)%billion, 1/1000)
	var s bytes.Buffer

	println(n)

	for n > 0 {
		switch {
		case n > hundred:
			s.WriteString(m[(n/100)*100])
			n %= hundred
			s.WriteString("-")
		case n > ten:
			s.WriteString(m[(n/10)*10])
			n %= ten
			s.WriteString("-")
		case n < ten:
			s.WriteString(m[n])
			n = 0

		}
	}

	return s.String(), true
}
