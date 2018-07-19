package main

import "fmt"

func main() {
	fmt.Println(
		func(i int) string{
		return fmt.Sprintf("i'm %v years old\n", i)
		}(31)		)
}
