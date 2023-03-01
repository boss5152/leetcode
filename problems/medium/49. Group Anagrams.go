package medium

import (
	"sort"
)

type Result struct {
	result []string
}

func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	map1 := make(map[string]*Result)

	for _, v := range strs {
		s := sortAlpha(v)
		if _, ok := map1[s]; !ok {
			map1[s] = &Result{}
			map1[s].result = append([]string{}, v)
		} else {
			map1[s].result = append(map1[s].result, v)
		}
	}

	for _, v := range map1 {
		result = append(result, v.result)
	}

	return result
}

// 字母排序
func sortAlpha(s string) string {
	strBytes := []byte(s)

	sort.Slice(strBytes, func(i, j int) bool {
		return strBytes[i] < strBytes[j]
	})

	return string(strBytes)
}
