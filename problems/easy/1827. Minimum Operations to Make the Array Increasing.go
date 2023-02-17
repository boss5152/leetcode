package easy

func MinOperations(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			result = result + nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}

	return result
}
