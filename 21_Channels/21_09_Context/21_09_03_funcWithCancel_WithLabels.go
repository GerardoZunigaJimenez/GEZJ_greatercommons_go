package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	fmt.Println("Is spiderman still alive?")
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("\tI don't wanna go mr stark")
				return // returning not to leak the goroutine
			case dst <- n:
				fmt.Printf("\tSpiderman is winning by %d time\n", n)
				n++
			}
		}
	}()
	return dst
}
