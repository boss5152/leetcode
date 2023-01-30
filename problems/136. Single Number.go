package problems

func SingleNumber(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums {
		if v, ok := m[num]; !ok {
			m[num] = struct{}{}
		} else if v == struct{}{} {
			delete(m, num)
		}
	}

	// only 1 element should be left
	for res := range m {
		return res
	}

	return 0 // never reach here
}
