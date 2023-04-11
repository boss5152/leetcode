package sort

import "fmt"

func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minId := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minId] > nums[j] {
				minId = j
			}
		}

		nums[minId], nums[i] = nums[i], nums[minId]
	}

	fmt.Println(nums)
}
