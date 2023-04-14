package matrix

type Pair struct {
	i int
	j int
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for i, r := range *m {
		for j, el := range r {
			saddleInRow := true
			for x, rowItem := range (*m)[i] {
				if x != j && rowItem > el {
					saddleInRow = false
					break
				}
			}
			if !saddleInRow {
				continue
			}
			saddleInCol := true
			for y, col := range *m {
				if y != i && col[j] < el {
					saddleInCol = false
					break
				}
			}
			if !saddleInCol {
				continue
			}

			pairs = append(pairs, Pair{i + 1, j + 1})
		}
	}
	return pairs
}
