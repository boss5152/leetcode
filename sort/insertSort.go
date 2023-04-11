package sort

import "fmt"

func InsertSort(nums []int) {
	nSort := []int{nums[0]}
	for v := range nums {
		for i := 0; i < len(nSort); i++ {
			if v < nSort[i] {
				nSort = append(append(nSort[:i], v), nSort[i:]...)
				break
			}
		}
	}

	fmt.Println(nSort)
}
