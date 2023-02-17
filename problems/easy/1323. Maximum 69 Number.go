package easy

import (
	"strconv"
)

func Maximum69Number(num int) int {
	sNum := strconv.Itoa(num)
	change := true
	result := ""
	for _, v := range sNum {
		if change && v == '6' {
			result += "9"
			change = false
		} else {
			result += string(v)
		}
	}

	ans, _ := strconv.Atoi(result)

	return ans
}
