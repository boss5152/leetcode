package easy

func SortArrayByParity(nums []int) []int {
	even := []int{}
	odd := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			odd = append(odd, num)
		} else {
			even = append(even, num)
		}
	}

	return append(odd, even...)
}
