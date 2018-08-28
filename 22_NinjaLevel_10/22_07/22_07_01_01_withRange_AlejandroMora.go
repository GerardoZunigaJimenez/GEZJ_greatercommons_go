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

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	gen(&wg, c)

	go channelReader(&wg, c)
	wg.Wait()
	fmt.Println("Imprimeee maldita")
}

func gen(wg *sync.WaitGroup, c chan<- int) {
	for i := 0; i < goRoutines; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for j := 0; j < iteraionByGoRoutine; j++ {
				c <- r.Intn(1000)
			}
			wg.Done()
		}(wg)
	}
}

func channelReader(c <-chan int) {
	i := 0

	for {
		select {
		case v := <-c:
			i++
			fmt.Println("iterarion #", i, " with Random Number:", v)
		}
	}

}
