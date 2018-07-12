package main

import (
	"fmt"
)

func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	//Assign the func execution to "x" variable that will invoke the function
	fmt.Println(x())

	//calling directly the function that return a function
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}
