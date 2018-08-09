package main

import "fmt"

var rounds int = 10

func main() {
	c := make(chan int)
	go func() {
		for i:= 0; i<rounds; i++{
		c <- i
		}
		close(c)
	}()

	for {
		v, ok := <-c
		if !ok{
			fmt.Println("We finish")
			return
		}
		fmt.Println("\tvalue:",v )
	}
}
