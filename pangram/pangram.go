package pangram

import (
	"regexp"
	"strings"
)

// Tried to do this with regex, but go regex doesnt support lookaheads so the ensuing regex would have the following form:
// regexp.MustCompile(`.*(a.*b.*c|a.*c.*b|b.*a.*c|b.*c.*a|c.*a.*b|c.*b.*a).*`)
// just for three letters, meaning that for 26 there would be, the product of 26 possible unions
// which is the least performant solution
// https://codegolf.stackexchange.com/questions/172809/regular-expression-to-find-a-pangram

var pattern = regexp.MustCompile(`[a-z]`)

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
