package examination

import "fmt"

func main3() {
	a := []int{1, 2, 3, 4, 5}
	t := a[3:4:4]

	fmt.Println(cap(t), t)
}

/**
A. 1, [4]
B. 5, [4]
C. 1, [5]
D. 5, [5]
E. panic
**/
