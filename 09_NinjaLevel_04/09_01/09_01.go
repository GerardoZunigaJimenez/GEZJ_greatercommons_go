package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}

	for p, i := range a {
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n", p, i, i)
	}
}
