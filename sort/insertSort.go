package sort

import "fmt"

func InsertSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j = j - 1

			fmt.Println(nums)
		}
	}

	fmt.Println(nums)
}
