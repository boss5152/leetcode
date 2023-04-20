package examination

import "fmt"

func main2() {
	a := []int{4: 44, 22, 33, 1: 11, 55}

	fmt.Println(cap(a), a[5])
}

/**
A. 5, 55
B. 7, 55
C. 7, 33
D. 7, 22
E. panic
**/
