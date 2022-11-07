package pangram

import (
	"regexp"
	"strings"
)

// https://stackoverflow.com/questions/32557299/javascript-pangram-regex
// var pattern = regexp.MustCompile(`.([a-z])+`)
var pattern = regexp.MustCompile(`[a-z]`)

// var pattern = regexp.MustCompile(`/([a-z])(?!.*\1)/g`)

// func IsPangram(input string) bool {
// 	trimmed := strings.ToLower(strings.Replace(input, " ", "", -1))
// 	println(trimmed)
// 	return pattern.MatchString(trimmed)
// }

func IsPangram(input string) bool {
	letterMap := map[rune]bool{}
	toLower := strings.ToLower(input)

	for _, letter := range toLower {
		if pattern.MatchString(string(letter)) {
			letterMap[letter] = true
		}
	}
	return len(letterMap) == 26
}
