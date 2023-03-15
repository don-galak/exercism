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
	40:       "forty",
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

	var s bytes.Buffer
	for n > 0 {
		switch {
		case n > billion:
			divByHundred := n / billion
			s.WriteString(m[divByHundred] + " " + m[billion])
			n %= billion
		case n > million:
			divByHundred := n / million
			s.WriteString(m[divByHundred] + " " + m[million])
			n %= million
		case n > thousand:
			if s.Len() > 0 {
				s.WriteString(" ")
			}
			divByHundred := n / thousand
			s.WriteString(m[divByHundred] + " " + m[thousand])
			n %= thousand
		case n > hundred:
			if s.Len() > 0 {
				s.WriteString(" ")
			}
			divByHundred := n / hundred
			s.WriteString(m[divByHundred] + " " + m[hundred])
			n %= hundred
		case n >= ten:
			if s.Len() > 0 {
				s.WriteString(" ")
			}
			s.WriteString(m[(n/ten)*ten])
			n %= ten
			s.WriteString("-")
		case n < ten:
			s.WriteString(m[n])
			n = 0
		}
	}

	return s.String(), true
}

// func doThing(s bytes.Buffer, n int64) {
// 	// orig := n
// 	for n > 0 {

// 	}
// }
