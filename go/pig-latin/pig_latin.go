package piglatin

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`^[aeiuo]+$`)

const ay = "ay"

func Sentence(sentence string) string {
	s := strings.Split(sentence, " ")

	outArr := []string{}
	for _, word := range s {
		outArr = append(outArr, toPigLatin(word))
	}
	return strings.Join(outArr, " ")
}

func toPigLatin(sentence string) string {
	if reg.MatchString(sentence[:1]) || sentence[:2] == "xr" || sentence[:2] == "yt" {
		sentence += ay
		return sentence
	}
	index := 0
	for i, letter := range sentence {
		index = i
		if letter == 'y' && index == 0 {
			index++
			return sentence[index:] + sentence[:index] + ay
		}
		if letter == 'y' || reg.MatchString(string(letter)) {
			return sentence[index:] + sentence[:index] + ay
		}
		if letter == 'q' && sentence[i+1] == 'u' {
			index += 2
			return sentence[index:] + sentence[:index] + ay
		}
	}

	return ""
}
