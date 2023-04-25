package problems

import "fmt"

type GameBlackList struct {
	HallID map[int]HallList
}

type HallList struct {
	TagIdList map[int]TagList
}

type TagList struct {
	GameKind map[int][]int
}

type CatS struct {
	C  []int
	C1 map[int]int
	C2 map[int][]int
	C3 TagList
}

func Try() {
	gbl := GameBlackList{
		HallID: map[int]HallList{
			6: {
				TagIdList: map[int]TagList{
					10062: {
						GameKind: map[int][]int{
							5: {60052, 60053},
						},
					},
				},
			},
		},
	}

	fmt.Println(gbl)

	tt := struct {
		Hello []int
		World map[int]int
		Dog   map[int][]int
		Cat   CatS
	}{
		Hello: []int{1, 2, 3},
		World: map[int]int{
			1: 1,
		},
		Dog: map[int][]int{
			1: {1, 2, 3},
		},
		Cat: CatS{
			C: []int{1, 2, 3},
			C1: map[int]int{
				1: 1,
			},
			C2: map[int][]int{
				1: {
					1,
				},
			},
			C3: TagList{
				GameKind: map[int][]int{
					5: {60052, 60053},
				},
			},
		},
	}

	fmt.Println(tt)
}
