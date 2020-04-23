package _3_02_HelloWorld

import "fmt"

func main() {

	fmt.Println("Hello World")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	func() {
		fmt.Println("and then we exited")
	}()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
