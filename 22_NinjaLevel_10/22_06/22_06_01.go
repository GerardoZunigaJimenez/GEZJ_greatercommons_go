package main

import (
	"math/rand"
	"fmt"
)

var iterations = 10

func main() {
	c := make(chan int)

	go randomGenerator(c)

	channelReader(c)

	fmt.Println("\nI don't wanna go Mr Stark!!!")
}

func randomGenerator(c chan<- int) {
	for i := 0; i < iterations; i++ {
		j := rand.Intn(1000)
		c <- j
	}
	close(c)
}

func channelReader(c <-chan int){
	i := 0
	for{
		i++
		v, ok := <-c
		if ok{
			fmt.Println("iterarion #", i, " with Random Number:", v)
		} else{
			return
		}
	}
}
