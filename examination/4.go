package examination

import "fmt"

func main4() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		<-ch1
	}()

	go func() {
		select {
		case ch1 <- 1:
		case ch2 <- 2:
		}
	}()

	fmt.Println(<-ch2)
}

/**
A. 1
B. 2
C. deadlock
D. nil
E. 其他
**/
