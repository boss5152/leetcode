package problems

import "fmt"

func Learn() {
	n := 5
	// 創建一個 n^n 陣列
	result := [][]int{}
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		result = append(result, append([]int{}, slice...))
	}

	result[2][3] = 2

	fmt.Println(result)
}
