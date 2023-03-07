package medium

func GenerateMatrix(n int) [][]int {
	// 創建一個 n^n 陣列
	result := [][]int{}
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		result = append(result, append([]int{}, slice...))
	}

	return result
}
