package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	case !(4 == 4):
		fmt.Println(false)
	}

}
