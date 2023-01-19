package problems

func ReverseString(s []byte) []byte {
	ans := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		ans = append(ans, s[i])
	}

	copy(s, ans)

	return s
}
