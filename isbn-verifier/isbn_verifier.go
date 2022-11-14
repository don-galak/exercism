package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnLen := len(isbn)
	if isbnLen != 10 {
		return false
	}

	result := 0

	for i, digit := range isbn {
		if toNum, err := strconv.Atoi(string(digit)); err != nil {
			if i == isbnLen-1 && digit == 'X' {
				result = result + 10*(isbnLen-i)
			} else {
				return false
			}
		} else {
			result = result + toNum*(isbnLen-i)
		}
	}

	return result%11 == 0
}
