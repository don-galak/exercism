package isbn

func IsValidISBN(isbn string) bool {
	result := 0
	n := 0
	for i := 0; i < len(isbn); i++ {
		if isbn[i] <= byte('9') && isbn[i] >= byte('0') {
			result += int(isbn[i]-byte('0')) * (10 - n)
			n++
		} else if isbn[i] == byte('X') && n == 9 {
			result += 10
			n++
		} else if isbn[i] == byte('-') && n < 10 {
			continue
		} else {
			return false
		}
	}
	return result%11 == 0 && n == 10
}
