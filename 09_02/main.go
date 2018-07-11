package main

import "fmt"

func main() {
	a := []int{10,20,30,40,50,60,70,80,90,100}

	for p,i:= range a{
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n",p,i,i)
	}
}
