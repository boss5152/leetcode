package medium

func Partition(s string) [][]string {
	result := [][]string{}
	partitionHelper(s, []string{}, &result)

	return result
}

func partitionHelper(s string, ans []string, result *[][]string) {
	// 一個個拆分比對
	for i := 1; i <= len(s); i++ {
		if palindrome(s[:i]) {
			partitionHelper(s[i:], append(ans, s[:i]), result)
		}
	}

	if len(s) == 0 {
		*result = append(*result, ans)
	}
}

// 判斷回文
func palindrome(s string) bool {
	// 單一個
	if len(s) == 1 {
		return true
	}

	// 回文比對
	// 偶數
	if len(s)%2 == 0 {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
	} else { // 奇數
		for i := 0; i <= len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
	}

	return true
}
