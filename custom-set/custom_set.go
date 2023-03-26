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

func (s Set) length() int {
	return len(s.toSlice())
}

func (s Set) IsEmpty() bool {
	return len(s) <= 2
}

func (s Set) Has(elem string) bool {
	return strings.Contains(s.String(), fmt.Sprintf(`"%s"`, elem))
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		elementToAdd := ""
		if s.length() > 0 {
			elementToAdd = fmt.Sprintf(`, "%s"}`, elem)
		} else {
			elementToAdd = fmt.Sprintf(`"%s"}`, elem)
		}
		*s = Set(strings.Replace(s.String(), "}", elementToAdd, 1))
	}
}

func (s Set) removeBraces() string {
	return s.String()[1 : len(s.String())-1]
}

func Subset(s1, s2 Set) bool {
	return strings.Contains(s2.String(), s1.removeBraces())
}

func (s Set) toSlice() []string {
	if s.IsEmpty() {
		return make([]string, 0)
	}

	sanitizedSet := strings.ReplaceAll(s.String()[1:len(s.String())-1], " ", "")
	return strings.Split(sanitizedSet, ",")
}

func Disjoint(s1, s2 Set) bool {
	slice1 := s1.toSlice()
	slice2 := s2.toSlice()

	for _, i := range slice1 {
		for _, j := range slice2 {
			if i == j {
				return false
			}
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if s1.length() != s2.length() {
		return false
	}

	slice1 := s1.toSlice()
	slice2 := s2.toSlice()

	for _, i := range slice1 {
		elementExists := false
		for _, j := range slice2 {
			if i == j {
				elementExists = true
			}
		}
		if !elementExists {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	intersection := []string{}

	slice1 := s1.toSlice()
	slice2 := s2.toSlice()

	for _, i := range slice1 {
		for _, j := range slice2 {
			if i == j {
				intersection = append(intersection, strings.ReplaceAll(i, `"`, ""))
			}
		}
	}
	return NewFromSlice(intersection)
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
