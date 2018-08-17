package main

import (
	"fmt"
)

var iterations = 10

func main() {
	q := make(chan int)
	c := make(chan int)
	go gen(c, q)

	receive(c, q)

	fmt.Println("about to exit")
}

// the action that is writing to the channel
// should to be a go routine to generate the blocking
func gen(c, q chan<- int) {
	for i := 0; i < iterations; i++ {
		c <- i
	}
	close(c)
	q <- 1
}

func receive(c, q <-chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			}
		case <-q:
			return
		}
	}
}
