package main

import "fmt"

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func subtraction(xi ...int) int {
	total := 0
	for _, v := range xi {
		total -= v
	}
	return total
}

func multiplication(xi ...int) int {
	total := 1
	for _, v := range xi {
		total *= v
	}
	return total
}

func main() {
	numbers := []int{45, 78, 1, 6, 87, 63, 15, 6, 74, 8, 6, 1, 5, 3, 5, 3, 45}

	oList := []func(...int) int{sum, subtraction, multiplication}

	fmt.Printf("Operation List Type: %T\n", oList)
	fmt.Println("Numbers that will be operated: ", numbers)

	fmt.Println()
	for _, function := range oList {
		fmt.Printf("The operation '%T' has the result of %v\n", function, function(numbers...))
	}

}
