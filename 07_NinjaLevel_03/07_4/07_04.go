package main

import "fmt"

func main() {
	bd := 1986
	for {
		if bd > 2018 {
			break
		}

		fmt.Println("years to be alived: ", bd)
		bd++
	}
}
