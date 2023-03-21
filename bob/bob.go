package bob

import (
	"regexp"
	"strings"
)

var upperCase = regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	isQuestion := remark[len(remark)-1] == '?'
	isUpperCase := upperCase.MatchString(remark)

	switch {
	case isQuestion && isUpperCase:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isUpperCase:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
