package main

import "fmt"

func main() {
	defer foo()
	bar()
}

// https://blog.golang.org/defer-panic-and-recover
func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}