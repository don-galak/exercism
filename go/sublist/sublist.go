package sublist

import (
	"bytes"
	"strconv"
	"strings"
)

func write(l []int) string {
	var s bytes.Buffer
	for _, x := range l {
		s.WriteString(strconv.Itoa(x) + ".")
	}
	return s.String()
}

func Sublist(l1, l2 []int) Relation {
	s1 := write(l1)
	s2 := write(l2)

	switch {
	case s1 == s2:
		return RelationEqual
	case strings.Contains(s2, s1):
		return RelationSublist
	case strings.Contains(s1, s2):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}
