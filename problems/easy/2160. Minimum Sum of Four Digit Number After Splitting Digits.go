package easy

import (
	"sort"
	"strconv"
)

func MinimumSum(num int) int {
	intSlice := []int{}

	// 轉string
	sNum := strconv.Itoa(num)

	for _, v := range sNum {
		i, _ := strconv.Atoi(string(v))

		intSlice = append(intSlice, i)
	}

	sort.Ints(intSlice)

	return intSlice[0]*10 + intSlice[1]*10 + intSlice[2] + intSlice[3]
}
