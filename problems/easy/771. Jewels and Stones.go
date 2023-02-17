package easy

func NumJewelsInStones(jewels string, stones string) int {
	jMap := map[rune]struct{}{}
	count := 0

	for _, v := range jewels {
		jMap[v] = struct{}{}
	}

	for _, v := range stones {
		if _, ok := jMap[v]; ok {
			count++
		}
	}

	return count
}
