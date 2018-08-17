package main

import (
	"fmt"
)

//The capacity, in number of elements, sets the size of the buffer in the channel.
// If the capacity is zero or absent, the channel is unbuffered and communication
// succeeds only when both a sender and receiver are ready.” Golang Spec

func main() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
