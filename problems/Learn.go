package problems

import (
	"fmt"
	"sync"
)

func Learn() {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i += 2 {
			<-ch
			fmt.Println(i)
			ch <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			ch <- struct{}{}
			<-ch
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
