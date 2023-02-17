package medium

import "math"

func LengthOfLongestSubstring(s string) int {
	sMap := make(map[string]struct{})
	lens := 0

	for _, v := range []byte(s) {
		if _, ok := sMap[string(v)]; ok {
			sMap = map[string]struct{}{}
		}

		sMap[string(v)] = struct{}{}

		if lens < len(sMap) {
			lens = len(sMap)
		}
	}

	if len(s) <= 1 {
		return len(s)
	}
	lo, hi, ans := 0, 0, 1.0
	m := make(map[byte]bool)

	for hi < len(s) {
		c := s[hi]
		hi++
		for lo < hi {
			if _, present := m[c]; !present {
				break
			}
			lowC := s[lo]
			lo++
			delete(m, lowC)
		}
		m[c] = true
		ans = math.Max(ans, float64(hi-lo))
	}
	return int(ans)
}
