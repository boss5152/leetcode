package easy

func RepeatedNTimes(nums []int) int {
	result, max := 0, 0
	map1 := map[int]int{}
	for _, v := range nums {
		if _, ok := map1[v]; !ok {
			map1[v] = 1
		} else {
			map1[v]++
		}

		if map1[v] > max {
			max = map1[v]
			result = v
		}
	}

	return result
}
