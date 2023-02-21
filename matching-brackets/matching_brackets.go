package brackets

func Bracket(input string) bool {
	startingSquare := false
	startingBracket := false
	startingParenthese := false
	// endingSquare := false
	// endingBracket := false
	// endingParenthese := false

	for _, c := range input {
		switch c {
		case '[':
			startingSquare = true
		case ']':
			if !startingSquare {
				return false
			}
			startingSquare = false
		case '{':
			startingBracket = true
		case '}':
			if !startingBracket {
				return false
			}
			startingBracket = false
		case '(':
			startingParenthese = true
		case ')':
			if !startingParenthese {
				return false
			}
			startingParenthese = false
		}
	}
	if startingBracket || startingSquare || startingParenthese {
		return false
	}

	return true
}
