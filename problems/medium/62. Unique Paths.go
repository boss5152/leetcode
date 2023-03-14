package medium

func UniquePaths(m int, n int) int {
	slice1 := []int{}
	result := [][]int{}

	// 建造field
	for i := 0; i < n; i++ {
		slice1 = append(slice1, 1)
	}

	for i := 0; i < m; i++ {
		result = append(result, append([]int{}, slice1...))
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	return result[m-1][n-1]
}
