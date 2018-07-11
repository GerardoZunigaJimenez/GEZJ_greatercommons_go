package main

import "fmt"

func main() {
	for i:= 33; i <= 1222; i++{
		fmt.Printf("%v\t%#x\t%#U\n",i, i, i)
	}
}
