package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	fanin := make(chan int)

	go send(even, odd, quit)

	go receive(even, odd, quit, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receive(even, odd, quit <-chan int, fanin chan<- int) {
thisLoop:
	for {
		select {
		case fanin <- <-even:
			fmt.Println("removed value from the even channel")
		case fanin <- <-odd:
			fmt.Println("removed value from the odd channel")
		case <-quit:
			break thisLoop
		}
	}
	close(fanin)
}
