package problems

import "fmt"

type First struct {
}

func foo(i interface{}) {

}

func boo(i *interface{}) {

}

func Try() {
	input := []int{20, 70, -40, 30, -10}
	// output := []int{20, 70, 30, -40, -10}

	slice1, slice2 := []int{}, []int{}
	for _, v := range input {
		if v > 0 {
			slice1 = append(slice1, v)
		} else {
			slice2 = append(slice2, v)
		}
	}

	fmt.Println(append(slice1, slice2...))
}
