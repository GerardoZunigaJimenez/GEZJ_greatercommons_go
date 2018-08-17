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
	c := make(chan string, iterations)
	q := make(chan bool, 1)
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
func gen(c chan<- string, q chan<- bool) {

	for i := 0; i < iterations; i++ {
		fmt.Println("\tInsertion #", i)
		c <- fmt.Sprint("value: ", i)
	}

	close(c)
	q <- true
}

func receive(c <-chan string, q <-chan bool) {
	go func() {
		for {
			select {
			// we could check the ok variable to know
			// if the value was inserted or it is garbage
			// from the channel
			case v, ok := <-c:
				if ok {
					fmt.Println(v)
				}
			case <-q:
				wg.Done()
			}
		}
	}()
}
