package main

import (
	"fmt"
)
//IMPORTANT: CHANNELS BLOCK
//channels allow
//coordination / synchronization / orchestration
//buffering (buffered channels)
func main() {
	ca := make(chan int)
	ca <- 42
	fmt.Println(<-ca)
}