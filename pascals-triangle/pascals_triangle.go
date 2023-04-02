package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	triangle[0] = []int{1}

	for i := 1; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}
