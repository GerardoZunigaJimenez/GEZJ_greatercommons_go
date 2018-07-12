package main

import "fmt"

func main() {
	x := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...)

	fmt.Println(x)
	fmt.Println(y)

	}
