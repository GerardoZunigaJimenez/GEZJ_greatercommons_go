package main

import "fmt"

func main() {
	fmt.Println("'&' gives you the address")
	fmt.Println("'*' gives you the value stored at an address when you have an address")

	a := 42
	fmt.Println("\nVAR")
	fmt.Println("\t'a' value:", a)
	fmt.Println("\t'a' address:", &a) // '&' gives you the address

	fmt.Printf("\t'a' value type: %T\n", a)
	fmt.Printf("\t'a' address from a var: %T\n", &a)
	fmt.Println("\t'a' deferencing address of a var: ", *&a)

	fmt.Println("\nPointer")
	var b *int = &a
	fmt.Println("\t'b' value: ", b)
	fmt.Printf("\t'b' type: %T\n", b)
	fmt.Println("\t'b' deferencing address: ", *b) // '*' gives you the value stored at an address when you have an address

	fmt.Println("\nPointer")
	c := &a
	fmt.Println("\t'c' addres: ", c)
	fmt.Printf("\t'c' type: %T\n", c)
	fmt.Printf("\t'c' deferencing address: %v\n", *c)

	fmt.Println("\nChanging the value of the variable '*b = 52' ")
	*b = 52
	fmt.Println("\t'a' = value: ", a)
}
