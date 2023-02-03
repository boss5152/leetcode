package problems

func Intersection349(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]struct{})
	ans := []int{}

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num2 == num1 {
				map1[num1] = struct{}{}
			}
		}
	}

	for k, _ := range map1 {
		ans = append(ans, k)
	}

	return ans
}
