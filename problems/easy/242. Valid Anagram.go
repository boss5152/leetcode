package easy

import (
	"reflect"
	"strings"
)

func IsAnagram(s string, t string) bool {
	var (
		// 存放不重複的字母用
		sTemp = map[string]struct{}{}
		tTemp = map[string]struct{}{}
		// 計數map
		sNum = make(map[string]int)
		tNum = make(map[string]int)
	)

	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	for _, v := range sSlice {
		if _, ok := sTemp[v]; !ok {
			sTemp[v] = struct{}{}
			sNum[v] = 1
		} else {
			sNum[v] += 1
		}
	}

	for _, v := range tSlice {
		if _, ok := tTemp[v]; !ok {
			tTemp[v] = struct{}{}
			tNum[v] = 1
		} else {
			tNum[v] += 1
		}
	}

	return reflect.DeepEqual(sNum, tNum)
}
