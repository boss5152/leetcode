package examination

import "fmt"

func main5() {
	var x interface{} = 2.0

	if y, ok := x.(int); ok {
		fmt.Println(y)
	} else {
		fmt.Println(int(y))
	}
}

/**
A. 2
B. 2.0
C. 0
D. panic
E. 其他
**/
