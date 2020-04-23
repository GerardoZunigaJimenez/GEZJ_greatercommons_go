package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(foo())

	fmt.Println(bar())
}

func foo() int {
	return rand.Int()
}

func bar() (int, string) {
	return rand.Int(), "Hello World"
}
