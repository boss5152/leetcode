package medium

import "fmt"

func PartitionLabels(s string) []int {
	var ret []int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	need := make(map[byte]int)
	for i := range s {
		need[s[i]] = i
	}

	for k, v := range need {
		fmt.Println(string(k), v)
	}

	start, end := 0, 0
	for i := range s {
		end = max(end, need[s[i]])
		if end == i {
			ret = append(ret, end-start+1)
			start = end + 1
		}
	}

	fmt.Println(ret)
	return ret
}
