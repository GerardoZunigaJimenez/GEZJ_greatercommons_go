package main

import "fmt"

func main() {
	x := []int{4,5,7,8,42}

	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	fmt.Println()
	for i, v:= range x {
		fmt.Println(i, v)
	}

	fmt.Println()
	for i := 0; i < 5; i++{
		fmt.Println(i,x[i])
	}
}