package problems

import (
	"strings"
)

func ReverseWords(s string) (ans string) {
	temp := strings.Split(s, " ")

	for k, v := range temp {
		vByte := []byte(v)
		reverse := []byte{}

		for i := len(vByte) - 1; i >= 0; i-- {
			reverse = append(reverse, vByte[i])
		}

		ans += string(reverse)

		if k != len(temp)-1 {
			ans += " "
		}
	}

	return ans
}
