package medium

func MaxProfit(prices []int) int {
	buy := -1
	sell := -1
	result := 0

	for i := 0; i < len(prices); i++ {
		// 最後一天
		if i+1 == len(prices) {
			if sell != -1 {
				result += sell - buy
			}

			return result
		}

		// 明天漲
		if prices[i+1] > prices[i] {
			if buy == -1 {
				buy = prices[i]
			}

			sell = prices[i+1]
		}

		// 明天跌
		if prices[i+1] < prices[i] {
			if sell != -1 {
				result += sell - buy
				buy, sell = -1, -1
			}
		}
	}

	return result
}
