package easy

func RemoveDuplicates(nums []int) int {
	repeat := []int{}
	for _, num := range nums {
		// 是否存在
		exist := false

		// 遍歷一遍，存在的話不放入repeat裡
		for _, rV := range repeat {
			if rV == num {
				exist = true
			}
		}

		if !exist {
			repeat = append(repeat, num)
		}
	}

	return len(nums)
}
