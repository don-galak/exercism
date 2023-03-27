package stringset

import (
	"bytes"
	"fmt"
)

type Set map[string]bool

func New() Set {
	return make(Set, 0)
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, e := range l {
		s.Add(e)
	}
	return s
}

func (s Set) String() string {
	sLen := len(s)
	var st bytes.Buffer
	st.WriteString("{")

	for e := range s {
		element := fmt.Sprintf(`"%s"`, e)
		st.WriteString(element)
		if sLen--; sLen > 0 {
			st.WriteString(", ")
		}
	}
	st.WriteByte('}')
	return st.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(e string) bool {
	_, ok := s[e]
	return ok
}

func (s Set) Add(e string) {
	s[e] = true
}

func Subset(s1, s2 Set) bool {
	for e := range s1 {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	intersection := make(Set, 0)
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	for e := range s1 {
		if s2.Has(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := make(Set)
	for e := range s1 {
		if !s2.Has(e) {
			difference.Add(e)
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	for e := range s2 {
		s1.Add(e)
	}
	return s1
}
