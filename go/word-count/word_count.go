package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`\w+('\w+)?`)

func WordCount(phrase string) Frequency {
	f := make(Frequency)
	words := re.FindAllString(phrase, -1)
	for _, w := range words {
		word := strings.ToLower(w)
		f[word] += 1
	}
	return f
}
