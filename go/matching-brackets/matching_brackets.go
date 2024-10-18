package brackets

var bMap = map[rune]rune{')': '(', '}': '{', ']': '['}

func Bracket(s string) bool {
	var stack []rune
	for _, b := range s {
		switch b {
		case '[', '{', '(':
			stack = append(stack, b)
		case ']', '}', ')':
			n := len(stack) - 1
			if n < 0 || stack[n] != bMap[b] {
				return false
			}
			stack = stack[:n]
		}
	}
	return len(stack) == 0
}
