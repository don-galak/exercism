package sublist

import (
	"bytes"
	"fmt"
	"strings"
)

func toString(l []int) string {
	var s bytes.Buffer
	for _, x := range l {
		s.WriteString(fmt.Sprint(x, "."))
	}
	return s.String()
}

func Sublist(l1, l2 []int) Relation {
	ls1 := toString(l1)
	ls2 := toString(l2)

	switch {
	case ls1 == ls2:
		return RelationEqual
	case strings.Contains(ls2, ls1):
		return RelationSublist
	case strings.Contains(ls1, ls2):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}
