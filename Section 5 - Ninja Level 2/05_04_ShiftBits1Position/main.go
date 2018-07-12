package main

import "fmt"

func main() {
	a := 42
	b := a << 1

	fmt.Printf("%d\t\t\t%b\t\t\t\t\t%#x\n", a, a, a)
	fmt.Printf("%d\t\t\t%b\t\t\t\t\t%#x\n", b, b, b)

}
