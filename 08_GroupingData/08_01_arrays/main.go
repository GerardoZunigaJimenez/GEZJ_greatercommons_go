package main

import "fmt"

func main() {
	var x [5]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}
	fmt.Println(x)
}
