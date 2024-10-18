package school

import "sort"

type Grade struct {
	index    int
	students []string
}

type School map[int]Grade

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	if _, ok := (*s)[grade]; ok {
		st := (*s)[grade].students
		st = append(st, student)

		(*s)[grade] = Grade{grade, st}
		return
	}
	(*s)[grade] = Grade{grade, []string{student}}
}

func (s *School) Grade(level int) []string {
	return (*s)[level].students
}

func (s *School) Enrollment() []Grade {
	g := make([]Grade, 0)

	for _, gr := range *s {
		sort.Strings(gr.students)
		g = append(g, gr)
	}

	sort.Slice(g, func(i, j int) bool {
		return g[i].index < g[j].index
	})

	return g
}
