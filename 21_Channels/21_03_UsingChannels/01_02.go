package main

import "fmt"

func main() {

	c := make(chan int)
	// send
	for i := 0; i < 10; i++ {
		go foo(c, i)
	}
	// receive
	bar(c)

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int, i int) {
	c <- i
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
