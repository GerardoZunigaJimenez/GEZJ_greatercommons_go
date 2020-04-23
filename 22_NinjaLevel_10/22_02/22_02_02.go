package main

import "fmt"

/*
https://play.golang.org/p/_DBRueImEq
*/
func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
