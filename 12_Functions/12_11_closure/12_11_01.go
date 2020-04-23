package main

import "fmt"

func main() {
	var x int
	fmt.Println("x", x)
	x++
	{
		y := 1
		fmt.Println("y", y)
	}

	fmt.Println("x", x)
	foo()
}

func foo() {
	fmt.Println("hello")

}
