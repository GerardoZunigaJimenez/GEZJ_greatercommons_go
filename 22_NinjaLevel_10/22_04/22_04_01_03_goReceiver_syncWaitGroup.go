package main

import (
	"fmt"
	"sync"
)

var iterations = 10


//Synchronization
var wg sync.WaitGroup
var mu sync.Mutex
var delta = 1

func main() {
	c := make(chan int, iterations )
	q := make(chan int, 1)
	wg.Add(1)

	gen(c, q)

	receive(c, q)

	wg.Wait()
	fmt.Println("about to exit")
}

// the action that is writing to the channel
// should to be a go routine to generate the blocking
// if we use the receiver like go func, we will have extra values
// on the channel because it is trying to recover
// more information
func gen(c, q chan<- int) {

	for i := 0; i < iterations; i++ {
		fmt.Println("\tInsertion #", i)
		c <- i
	}

	close(c)
	q <- 1
}

func receive(c, q <-chan int) {
	go func() {
		for {
			select {
			case v, ok := <-c:
				if ok {
					//reading more values than the sent elements
					fmt.Println(v)
				}
			case <-q:
				wg.Done()
			}
		}
	}()
}
