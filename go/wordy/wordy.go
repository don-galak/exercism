package wordy

import (
	"strconv"
	"strings"
)

var ops = map[string]string{"plus": "", "minus": "", "multiplied": "", "divided": ""}

func Answer(question string) (int, bool) {
	question = strings.TrimSuffix(question, "?")

	words := strings.Split(question, " ")
	_, err := strconv.Atoi(words[len(words)-1])
	if err != nil {
		return 0, false
	}

	opCount := 0
	numCount := 0
	s := []string{}
	for _, w := range words {
		if _, err := strconv.Atoi(w); err == nil {
			numCount++
			s = append(s, w)
			continue
		}
		if _, exists := ops[w]; exists {
			opCount++
			s = append(s, w)
		}
	}
	_, err = strconv.Atoi(s[0])
	if err != nil || numCount-opCount != 1 {
		return 0, false
	}

	operation := ""
	result := 0

	for _, w := range s {
		n, err := strconv.Atoi(w)
		if err != nil {
			operation = w
		} else {
			switch operation {
			case "plus":
				result += n
			case "minus":
				result -= n
			case "multiplied":
				result *= n
			case "divided":
				result /= n
			default:
				result = n
			}
		}
	}

	return result, true
}
