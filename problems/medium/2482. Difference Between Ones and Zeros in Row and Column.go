package medium

func OnesMinusZeros(grid [][]int) [][]int {
	colSum := make([]int, len(grid))
	rowSum := make([]int, len(grid[0]))
	diff := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		diff[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				colSum[i]--
				rowSum[j]--
			} else {
				colSum[i]++
				rowSum[j]++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			diff[i][j] = colSum[i] + rowSum[j]
		}
	}

	return diff
}
