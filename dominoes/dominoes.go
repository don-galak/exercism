package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 || (len(input) == 1 && input[0][0] == input[0][1]) {
		return input, true
	}
	currentDomino := input[0]
	for index, domino := range input {
		if index == 0 {
			continue
		}
		if domino[1] == currentDomino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		if domino[0] == currentDomino[1] {
			newDomino := Domino{currentDomino[0], domino[1]}
			newDominoes := append([]Domino{newDomino}, input[1:index]...)
			newDominoes = append(newDominoes, input[index+1:]...)
			oldSolution, ok := MakeChain(newDominoes)
			if ok {
				newSolution := append([]Domino{currentDomino, domino}, oldSolution[1:]...)
				return newSolution, true
			}
		}
	}
	return nil, false
}
