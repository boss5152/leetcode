package problems

func JudgeCircle(moves string) bool {
	Px, Py := 0, 0

	for _, v := range moves {
		switch v {
		case 'R':
			Px += 1
		case 'L':
			Px -= 1
		case 'U':
			Py += 1
		case 'D':
			Py -= 1
		}
	}

	return Px == 0 && Py == 0
}
