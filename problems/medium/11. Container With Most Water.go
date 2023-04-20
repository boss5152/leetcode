package medium

func MaxArea(height []int) int {
	result := 0
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			max := maxHelper(i, j, height[i], height[j])
			if result < max {
				result = max
			}
		}
	}

	return result
}

func maxHelper(wA, wB, hA, hB int) int {
	var wid, hei int
	if wA > wB {
		wid = wA - wB
	} else {
		wid = wB - wA
	}

	if hA > hB {
		hei = hB
	} else {
		hei = hA
	}

	return hei * wid
}
