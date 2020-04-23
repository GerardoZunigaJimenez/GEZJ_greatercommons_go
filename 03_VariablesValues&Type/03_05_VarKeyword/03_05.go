package main

import "fmt"

var j = 6
var z int

// scope
//where a variable exists and is accessible
//best practice: keep scope as “narrow” as possible
//non declaration  statement outside function body
//assign the zero value when a new variables is declared without initialization

func main(){
	x := 45
	fmt.Println(x)

	var y = 3
	fmt.Println(y)

	fmt.Println(j)

	foo()
}

func foo(){
	fmt.Println("again",j)
	fmt.Println("ooops",z)


}
