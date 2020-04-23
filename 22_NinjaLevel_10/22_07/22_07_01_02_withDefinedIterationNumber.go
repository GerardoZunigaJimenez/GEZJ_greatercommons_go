package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var r = rand.New(source)
var goRoutines = 10
var iteraionByGoRoutine = 10

func main() {

	c := make(chan int)
	gen(c)

	channelReader(c)
	fmt.Println("I don't wanna go Mr Stark")
}

func gen(c chan<- int) {
	for i := 0; i < goRoutines; i++ {
		go func() {
			for j := 0; j < iteraionByGoRoutine; j++ {
				c <- r.Intn(1000)
			}
		}()
	}
}

func channelReader(c <-chan int) {
	for i := 0; i < (goRoutines * iteraionByGoRoutine); i++ {
		fmt.Println("iterarion #", i, " with Random Number:", <-c)
	}
}
