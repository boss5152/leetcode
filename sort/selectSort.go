package sort

import "fmt"

func SelectSort(nums []int) {
	var min int
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}

	fmt.Println(nums)
}
