package bob

import (
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	isQuestion := strings.HasSuffix(remark, "?")
	upperCase := strings.ToUpper(remark)
	lowerCase := strings.ToLower(remark)
	isUpperCase := upperCase == remark && lowerCase != remark

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
