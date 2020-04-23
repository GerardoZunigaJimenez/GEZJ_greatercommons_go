package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	//The code is taking the slice and sending value by value like arguments
	fmt.Println("The total sum is:", sum("James", xi...))
	fmt.Println()
	fmt.Println("The total for 2nd sum is:", sum("James"))

}

//The variaditic parameter should to be last one
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println("Len", len(x))
	fmt.Println("cap", cap(x))

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
