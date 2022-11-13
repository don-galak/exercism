package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`[ |,|\n|\:|\!|\@|\&|\$|\%|\^|\.]`)

func WordCount(phrase string) Frequency {
	f := Frequency{}
	sanitized := re.Split(phrase, -1)
	for _, w := range sanitized {
		trimmedWord := strings.Trim(w, "' ")
		if len(trimmedWord) > 0 {
			f[strings.ToLower(trimmedWord)]++
		}
	}
	return f
}
