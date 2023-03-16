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

var numMap = map[int64]string{
	0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine", ten: "ten",
	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
	80: "eighty", 90: "ninety", hundred: "hundred", thousand: "thousand", million: "million",
	billion: "billion",
}

func Say(n int64) (string, bool) {
	if n > max || n < 0 {
		return "", false
	}
	if s, exists := numMap[n]; exists {
		if n >= hundred {
			return "one " + s, true
		}
		return s, true
	}

	var s bytes.Buffer
	for n > 0 {
		switch {
		case n > billion:
			doThing(&s, billion, &n)
		case n > million:
			doThing(&s, million, &n)
		case n > thousand:
			doThing(&s, thousand, &n)
		case n > hundred:
			doThing(&s, hundred, &n)
		case n > ten:
			doThing(&s, ten, &n, func() { s.WriteString(numMap[(n/ten)*ten]) })
			s.WriteString("-")
		case n == ten:
			doThing(&s, ten, &n, func() { s.WriteString(numMap[(n/ten)*ten]) })
		case n < ten:
			s.WriteString(numMap[n])
			n = 0
		}
	}

	return s.String(), true
}

func doThing(s *bytes.Buffer, k int64, n *int64, cbs ...func()) {
	if s.Len() > 0 {
		s.WriteString(" ")
	}
	divBy := *n / k
	if len(cbs) > 0 {
		cbs[0]()
	} else {
		s.WriteString(numMap[divBy] + " " + numMap[k])
	}
	*n %= k
}
