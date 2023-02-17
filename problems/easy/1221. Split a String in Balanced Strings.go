package easy

func BalancedStringSplit(s string) int {
	balance := 0
	result := 0
	for _, v := range s {
		if v == 'R' {
			balance++
		} else {
			balance--
		}

		if balance == 0 {
			result++
		}
	}

	return result
}
