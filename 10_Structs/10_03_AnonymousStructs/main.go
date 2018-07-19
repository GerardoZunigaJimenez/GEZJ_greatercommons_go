package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		"James",
		"Bond",
		32,
	}

	fmt.Println(p1)
}
