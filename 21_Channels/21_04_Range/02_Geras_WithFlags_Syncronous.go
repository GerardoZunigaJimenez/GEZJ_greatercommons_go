package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup
var mu sync.Mutex
var counter int = 10
var delta int = counter * 2

func main() {
	wg.Add(delta)

	c := make(chan int)

	for i := 0; i < counter; i++ {
		go receive(c)
	}

	for i := 0; i < counter; i++ {
		go send(c, i)
	}

	wg.Wait()
	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int, i int) {
	mu.Lock()
	runtime.Gosched()
	fmt.Printf("put: %v \tGoRoutines: %v\n", i, runtime.NumGoroutine())

	c <- i

	mu.Unlock()
	wg.Done()
}

// receive channel
func receive(c <-chan int) {
	//mu.Lock()
	//runtime.Gosched()
	fmt.Printf("\tget: %v \tGoRoutines: %v\n", <-c, runtime.NumGoroutine())
	//mu.Unlock()
	wg.Done()
}
