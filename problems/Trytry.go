package problems

import (
	"fmt"
)

func Learn() {
	slice1 := []string{"a", "b", "c", "d", "e"}

	slice1 = append(slice1[:1], slice1[1+1:]...)

	fmt.Println(slice1)
}
