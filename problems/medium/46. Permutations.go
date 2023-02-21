package medium

func Permute(nums []int) [][]int {
	result := [][]int{}

	permuteHelper(nums, []int{}, &result)

	return result
}

func permuteHelper(nums []int, nowStep []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, append([]int{}, nowStep...))
		return
	}

	for k, v := range nums {
		permuteHelper(
			append(
				append([]int{}, nums[:k]...),
				nums[k+1:]...,
			),
			append(nowStep, v),
			result,
		)
	}
}
