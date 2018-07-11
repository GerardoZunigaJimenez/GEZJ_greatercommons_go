package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++
	{
		y:= 1
		fmt.Println(y)
	}

	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")

}
