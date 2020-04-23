package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Put:", i)
		c <- i
	}
	close(c)
}

// receive channel
func receive(c <-chan int) {
	for v := range c {
		fmt.Println("\tGet:", v)
	}
}
