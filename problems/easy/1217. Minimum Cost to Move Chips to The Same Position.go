package easy

func MinCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			odd++
		} else {
			even++
		}
	}

	if odd > even {
		return even
	}

	return odd
}
