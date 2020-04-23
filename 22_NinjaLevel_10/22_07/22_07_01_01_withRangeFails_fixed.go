package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var r = rand.New(source)
var goRoutines = 10
var iteraionByGoRoutine = 10

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	gen(c)

	//The reader should be a go routine to allow the right channel use
	go channelReader(c)

	wg.Wait()
	close(c)

	fmt.Println("I don't wanna go Mr Stark")
}

func gen(c chan<- int) {
	for i := 0; i < goRoutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < iteraionByGoRoutine; j++ {
				c <- r.Intn(1000)
			}
			defer wg.Done()
		}()

	}
}

func channelReader(c <-chan int) {
	i := 0
	for v := range c {
		i++
		fmt.Println("iterarion #", i, " with Random Number:", v)
	}
}
