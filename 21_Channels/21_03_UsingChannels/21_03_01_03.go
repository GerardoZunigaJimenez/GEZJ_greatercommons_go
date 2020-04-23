package main

import "fmt"

func main() {

	c := make(chan int)
	// send
	for i := 0; i < 10; i++ {
		go foo(c, i)
	}
	// receive
	for i := 0; i < 10; i++ {
		bar(c)
	}
	fmt.Println("about to exit")
}

//send
func foo(c chan<- int, i int) {
	fmt.Println("put: ", i)
	c <- i
}

//receive
func bar(c <-chan int) {
	fmt.Println("\tget: ", <-c)

}
