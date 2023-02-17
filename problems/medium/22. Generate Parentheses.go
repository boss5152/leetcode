package medium

import (
	"fmt"
)

var result *[][]string

func GenerateParenthesis(n int) []string {
	sample := []string{}

	for i := 0; i < n*2; i++ {
		sample = append(sample, ")")
	}
	sample[0] = "("

	a := 1

	// generateParenthesisHelp(1, 1, sample, n*2, *result)

	// s1 := make([]string, n*2)
	// s2 := make([]string, n*2)
	// // 第一個左括弧一定在第一位
	// sample[0] = "("
	// for p := 1; p < n*2-1; p++ {
	// 	// 第一層，放入第二個左括弧
	// 	if p-1 > 1 { // 一層for只放一個左括弧，因此若扣掉起始的一個左括弧仍大於一個空格(空格預設都是右括弧)，代表錯誤break
	// 		break
	// 	}

	// 	copy(s1, sample)
	// 	s1[p] = "("
	// 	for q := p + 1; q < n*2-1; q++ {
	// 		// 第二層，放入第三個左括弧
	// 		if q-2 > 2 { // 一層for只放一個左括弧，因此若扣掉起始的一個左括弧與第一層的的左括弧，仍大於兩個空格(空格預設都是右括弧)，代表錯誤break
	// 			break
	// 		}

	// 		copy(s2, s1)
	// 		s2[q] = "("

	// 		*result = append(*result, append([]string{}, s2...))
	// 	}
	// }

	fmt.Println(result)

	return []string{}
}

func generateParenthesisHelp(index int, layer int, step []string, len int, result *[][]string) {
	for i := index; i < len-1; i++ {
		sCopy := make([]string, len)
		if index-layer > layer {
			copy(sCopy, step)
			sCopy[index] = "("

			if layer == len/2-1 {
				*result = append(*result, append([]string{}, sCopy...))
			}

			generateParenthesisHelp(i, layer+1, sCopy, len, result)
		}
	}
}
