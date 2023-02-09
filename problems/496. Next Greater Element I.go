package problems

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := []int{}
	same := false

	for q, num := range nums1 {
		for i := 0; i < len(nums2); i++ {
			if nums2[i] > num && same {
				ans = append(ans, nums2[i])
				same = false
				break
			}

			if num == nums2[i] { // 有發現一樣找下一個大的
				same = true
			}
		}

		// len小於當前遍歷次數，代表沒找到下一個大的，填入-1
		if len(ans) < q+1 {
			ans = append(ans, -1)
			same = false
		}
	}

	return ans
}
