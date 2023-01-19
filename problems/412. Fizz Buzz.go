package problems

import "strconv"

func FizzBuzz(n int) []string {
	ans := []string{}

	for i := 1; i <= n; i++ {
		temp := ""
		if i%3 == 0 {
			temp += "Fizz"
		}

		if i%5 == 0 {
			temp += "Buzz"
		}

		if temp == "" {
			temp = strconv.Itoa(i)
		}

		ans = append(ans, temp)
	}

	return ans
}
