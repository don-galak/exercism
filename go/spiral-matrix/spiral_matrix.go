package spiralmatrix

func SpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	r, c := 0, 0
	dr, dc := 0, 1
	for i := 1; i <= n*n; i++ {
		matrix[r][c] = i

		if r+dr < 0 || r+dr >= n || c+dc < 0 || c+dc >= n || matrix[r+dr][c+dc] != 0 {
			dr, dc = dc, -dr
		}
		r, c = r+dr, c+dc
	}

	return matrix
}
