package main

import "fmt"

var y = 42
//declared the variable with the identifier "z"
//is of type string
// and I assign the value "shaken, not stirred
//this is a static programming language
//a variable is DECLARED to hold a value of a certain TYPE
// not a DYNAMIC
var z = "Shaken, not stirred"

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n",y)

	fmt.Println(z)
	fmt.Printf("%T\n",z)
	//z = 43
}
