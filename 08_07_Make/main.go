package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println()

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println()

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println()

	for i := 0; i < 100; i++ {
		x = append(x, i+3423)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
