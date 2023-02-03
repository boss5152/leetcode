package problems

import "sort"

func Intersection2248(nums [][]int) []int {
	map1 := make(map[int]int)
	ans := []int{}
	len := len(nums)
	for _, numSlice := range nums {

		for _, num := range numSlice {
			_, exist := map1[num]

			if exist { // 每重複一次加 1
				map1[num] += 1
			} else { // 不存在初始化
				map1[num] = 1
			}

			if map1[num] == len { // 完全重複
				ans = append(ans, num)
			}
		}
	}

	sort.Ints(ans)

	return ans
}
