package main

import "fmt"

func main() {
	foo()
	bar("James")
	fmt.Println( woo("Moneypenny") )
	fmt.Println( mouse("Ian", "Fleming"))
}

//Defying a function
// func ( r receiver) identifier (parameters) return(s) { }
// we define our func with parameters
//We call our func with arguments
func foo() {
	fmt.Println("hello from foo")
}

//Passing Parameters to the function
//everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

//One Single Return
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ",s)
}

//Multiple Returns
// the return parameters should to have the parentesis
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello"`)
	b := false
	return a,b
}
