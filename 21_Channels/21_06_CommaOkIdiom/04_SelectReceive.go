package main

import "fmt"

var count int = 50

func foo(f, b chan int) {
	for i := 0; i < count * 2; i++ {
		select {
		case v := <-f:
			fmt.Println("from foo:", v)
		case v := <-b:
			fmt.Println("from bar:", v)
		}
	}
	fmt.Println("about to exit")
}

func main() {
	f := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			f <- i
		}
	}()
	go func() {
		for i := 0; i < count; i++ {
			b <- i
		}
	}()
	foo(f, b)
}
