package stringset

import (
	"bytes"
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set string

func New() Set {
	return "{}"
}

func NewFromSlice(l []string) Set {
	if l == nil {
		return New()
	}

	var s bytes.Buffer
	s.WriteString("{")
	for i, n := range l {
		element := fmt.Sprintf(`"%s"`, n)

		if strings.Contains(s.String(), element) {
			break
		}

		if i > 0 {
			s.WriteString(", ")
		}

		s.WriteString(element)
	}
	s.WriteString("}")
	return Set(s.String())
}

func (s Set) String() string {
	return string(s)
}

func (s Set) IsEmpty() bool {
	return len(s) <= 2
}

func (s Set) Has(elem string) bool {
	return strings.Contains(s.String(), fmt.Sprintf(`"%s"`, elem))
}

func (s Set) Add(elem string) {
	panic("Please implement the Add function")
}

func Subset(s1, s2 Set) bool {
	return strings.Contains(s1.String(), s2.String())
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	panic("Please implement the Equal function")
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
