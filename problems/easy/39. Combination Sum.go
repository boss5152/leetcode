package easy

import "sort"

func CombinationSum(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	combinationSumHelper(nums, []int{}, target, 0, 0, &result)
	return result
}

func combinationSumHelper(nums, subset []int, target, startIndex, sum int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := startIndex; i < len(nums) && sum+nums[i] <= target; i++ {
		combinationSumHelper(nums, append(subset, nums[i]), target, i, sum+nums[i], result)
	}
}
