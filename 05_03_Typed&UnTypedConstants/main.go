package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d int = 5
	e string = "constant"
	f float64 = 65.5
)


func main() {

	fmt.Printf("%v\t\t\t%T\n",a,a)
	fmt.Printf("%v\t\t\t%T\n",b,b)
	fmt.Printf("%v\t\t\t%T\n",c,c)

	fmt .Println()

	fmt.Printf("%v\t\t\t%T\n",d,d)
	fmt.Printf("%v\t\t\t%T\n",e,e)
	fmt.Printf("%v\t\t\t%T\n",f,f)



}
