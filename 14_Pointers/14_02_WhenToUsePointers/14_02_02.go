package main

import "fmt"

func main() {

	//Pointers allow you to share a value stored in some memory location. Use pointers when
	//    you don’t want to pass around a lot of data
	//    you want to change the data at a location
	//Everything in Go is pass by value. Drop any phrases and concepts you may have from other languages.
	// Pass by reference, pass by copy - forget those phrases.
	// “Pass by value.” That is the only phrase you need to know and remember.
	// That is the only phrase you should use. Pass by value. Everything in Go is pass by value.
	// In Go, what you see is what you get (wysiwyg). Look at what is happening. That is what you get.

	x := 0
	fmt.Println("'x' value: ", x)
	fmt.Println("'x' address: ", &x)
	fmt.Println()

	foo(&x)

	// based on the copy by value, x will never change its value
	fmt.Println()
	fmt.Println("'x' value: ", x)
	fmt.Println("'x' address: ", &x)
}

func foo(y *int) {
	fmt.Println("'y' address: ", &y)
	fmt.Println("'y' pointer address: ", y)
	fmt.Println("'y' pointer value: ", *y)
	*y = 43
	fmt.Println("'y' pointer value: ", *y)
}
