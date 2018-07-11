package main

import "fmt"

func main() {
	fmt.Println(4*3*2*1)

	n := factorial(4)
	fmt.Println(n)

	n2 := factorialLoop(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial (n - 1)
}

func factorialLoop(n int) int{
	result := 1
	for ; n>0; n--	{
		result = result * n
	}
	return result
}