package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("b", b())
	fmt.Println("b", b())

	fmt.Println()
	fmt.Println("a", a())
	fmt.Println("b", b())

	fmt.Println()
	fmt.Println("c", incrementor()())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
