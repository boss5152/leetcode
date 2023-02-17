package easy

func MajorityElement(nums []int) int {
	countMap := map[int]int{}
	max := 0
	ans := 0
	for _, num := range nums {
		countMap[num]++
	}

	for k, v := range countMap {
		if v > max {
			max = v
			ans = k
		}
	}

	return ans
}
