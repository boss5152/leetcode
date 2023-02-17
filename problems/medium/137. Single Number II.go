package medium

func SingleNumber2(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		// 計算次數
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	for k, v := range m {
		if v != 3 { // 非三次的return
			return k
		}
	}

	return 0 // 沒吻合的
}
