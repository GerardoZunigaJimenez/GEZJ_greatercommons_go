package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- int) {
	for i := 0; i < 100; i++ {
		/* Adding negative numbers when i is odd */
		var value = i
		for c := 0; c < i; c++ {
			if i%2 != 0 {
				value = value * -1
			}
		}

		if value > 0 {
			if i%2 == 0 {
				even <- value
			} else {
				odd <- value
			}
		} else {
			quit <- value
		}
	}
	close(quit)
}

// receive channel
func receive(even, odd <-chan int, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("\tFinish comma ok bit", i)
				return
			} else {
				fmt.Println("\tFrom comma ok bit", i)
			}
		}
	}
}
