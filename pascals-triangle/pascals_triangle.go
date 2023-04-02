package pascal

func Triangle(n int) [][]int {
	firstRow := []int{1}
	triangle := make([][]int, n)
	triangle[0] = firstRow

	for i := 1; i < n; i++ {
		row := make([]int, i+1)

		for j := 0; j < i+1; j++ {
			prevRow := triangle[i-1]

			if j-1 < 0 {
				row[j] = prevRow[j]
			} else if len(prevRow) == j {
				row[j] = prevRow[j-1]
			} else {
				row[j] = prevRow[j-1] + prevRow[j]
			}
		}
		triangle[i] = row
	}

	return triangle
}
