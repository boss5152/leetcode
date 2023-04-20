package easy

func RemoveDuplicates(nums []int) int {
	map1 := make(map[int]bool)
	result := 0

	for _, k := range nums {
		if _, ok := map1[k]; !ok {
			map1[k] = true
			result++
		}
	}

	return result

	// repeat := []int{}
	// for _, num := range nums {
	// 	// 是否存在
	// 	exist := false

	// 	// 遍歷一遍，存在的話不放入repeat裡
	// 	for _, rV := range repeat {
	// 		if rV == num {
	// 			exist = true
	// 		}
	// 	}

	// 	if !exist {
	// 		repeat = append(repeat, num)
	// 	}
	// }

	// return len(nums)
}
