package medium

import (
	"math"
)

func IntToRoman(num int) string {
	result := ""
	romanMap := [][]string{
		{"M", "", ""},
		{"C", "D", "M"},
		{"X", "L", "C"},
		{"I", "V", "X"},
	}
	sliceInt := []int{}

	for i := 3; i >= 0; i-- {
		remainder := num / int(math.Pow(10, float64(i)))
		num = num % int(math.Pow(10, float64(i)))
		sliceInt = append(sliceInt, remainder)
	}

	for k, v := range sliceInt {
		result += intToRomanHelper(v, romanMap[k][0], romanMap[k][1], romanMap[k][2])
	}

	return result
}

func intToRomanHelper(num int, one, five, ten string) string {
	roman := ""
	switch num {
	case 1, 2, 3:
		for i := 1; i <= num; i++ {
			roman += one
		}
	case 4:
		roman = one + five
	case 5:
		roman = five
	case 6, 7, 8:
		roman += five
		for i := 6; i <= num; i++ {
			roman += one
		}
	case 9:
		roman = one + ten
	}

	return roman
}
