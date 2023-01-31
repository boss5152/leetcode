package problems

func CountBits(n int) []int {
	ans := []int{}

	for i := 0; i <= n; i++ {
		ans = append(ans, addBinary(i))
	}

	return ans
}

// 轉二進位相加
func addBinary(n int) (sum int) {
	if n == 0 {
		return 0
	}

	// 加餘數
	sum += n % 2

	if n/2 > 1 { // 大於1繼續
		sum += addBinary(n / 2)
	} else { // 不大於1加至總和
		sum += n / 2
	}

	return sum
}
