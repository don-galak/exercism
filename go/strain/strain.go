package strain

type Ints []int
type Lists [][]int
type Strings []string

func genericFilter[Es ~[]E, E int | []int | string](elements Es, filter func(E) bool) (out Es) {
	for _, element := range elements {
		if filter(element) {
			out = append(out, element)
		}
	}
	return
}

func (i Ints) Keep(filter func(int) bool) (out Ints) {
	return genericFilter(i, filter)
}

func (i Ints) Discard(filter func(int) bool) (out Ints) {
	return i.Keep(func(num int) bool {
		return !filter(num)
	})
}

func (l Lists) Keep(filter func([]int) bool) (out Lists) {
	return genericFilter(l, filter)
}

func (s Strings) Keep(filter func(string) bool) (out Strings) {
	return genericFilter(s, filter)
}
