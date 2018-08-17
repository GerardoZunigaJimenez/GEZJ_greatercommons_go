package main

import "fmt"

/*
Tests must

be in a file that ends with “_test.go”
put the file in the same package as the one being tested
be in a func with an signature “func TestXxx(*testing.T)”
Run a test

go test
Deal with test failure

use t.Error to signal failure
Nice idiom

expected
got

https://godoc.org/testing
http://www.golang-book.com/books/intro/12
 */
func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 9 =", mySum(5, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
