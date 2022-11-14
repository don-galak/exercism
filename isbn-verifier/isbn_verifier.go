package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnLen := len(isbn)
	if isbnLen != 10 {
		return false
	}
	result := 0

	for i, digit := range isbn {
		if unicode.IsDigit(digit) {
			result = result + int(digit-'0')*(isbnLen-i)
		} else if i == isbnLen-1 && digit == 'X' {
			result = result + 10*(isbnLen-i)
		} else {
			return false
		}
	}

	return result%11 == 0
}
