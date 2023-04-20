package problems

import (
	"fmt"
	"os"
	"sync"
)

var numForLearn = 10
var once sync.Once

func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func Try() {
	a := []int{4: 44, 22, 33, 1: 11, 55}

	fmt.Println(cap(a), a[5])
	os.Exit(1)
	// a channel of data type channel of data type string
	// 建立一個 channel 可以讀寫另一個（可以讀寫 string）的 channel
	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c)
	c <- "John"

	fmt.Println("main() stopped")
}

// main() started
// Hello John!
// main() stopped

func chM() {
	ch1, ch2 := make(chan int, numForLearn), make(chan int, numForLearn)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go f1(ch1, wg)
	go f2(ch1, ch2, wg)
	wg.Wait()

	for v := range ch2 {
		fmt.Println(v)
	}
}

func f1(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numForLearn; i++ {
		fmt.Println("1")
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if v, ok := <-ch1; ok {
			fmt.Println("2")
			ch2 <- v * v
		} else {
			break
		}
	}

	once.Do(func() { close(ch2) })
}
