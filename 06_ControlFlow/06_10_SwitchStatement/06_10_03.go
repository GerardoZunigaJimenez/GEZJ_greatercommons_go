package main

import "fmt"

func main() {
	n := "Bond"
	switch n{
	case "MoneyPennny", "Bond", "Dr no":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
