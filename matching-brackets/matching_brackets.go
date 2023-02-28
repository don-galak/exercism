package brackets

var bMap = map[rune]rune{')': '(', '}': '{', ']': '['}

func Bracket(s string) bool {
	var stack []rune
	for _, b := range s {
		switch b {
		case '[', '{', '(':
			stack = append(stack, b)
		case ']', '}', ')':
			if len(stack) == 0 || stack[len(stack)-1] != bMap[b] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
