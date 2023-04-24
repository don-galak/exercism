package school

type Grade struct {
	index    int
	students []string
}

type School struct {
	grades map[int]Grade
}

func New() *School {
	return &School{make(map[int]Grade)}
}

func (s *School) Add(student string, grade int) {
	if _, ok := s.grades[grade]; ok {
		st := s.grades[grade].students
		st = append(st, student)

		s.grades[grade] = Grade{grade, st}
		return
	}
	s.grades[grade] = Grade{grade, []string{student}}
}

func (s *School) Grade(level int) []string {
	return s.grades[level].students
}

func (s *School) Enrollment() []Grade {
	g := make([]Grade, 0)

	for _, gr := range s.grades {
		g = append(g, gr)
	}

	return g
}
