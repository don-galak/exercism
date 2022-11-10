package anagram

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	toArr := strings.Split(strings.ToLower(str), "")
	sort.Strings(toArr)
	return strings.Join(toArr, "")
}

func Detect(subject string, candidates []string) (out []string) {
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		if sortString(subject) == sortString(candidate) {
			out = append(out, candidate)
		}
	}

	return
}
