package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

//send
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	fmt.Printf("Printing Channel Content from: %T\n", c)
	i := 0
	for v := range c {
		if i == 10 {
			i = 0
			fmt.Printf("\n")

		}
		fmt.Printf("\t%v", v)
		i++
	}
	fmt.Printf("\n")
}
