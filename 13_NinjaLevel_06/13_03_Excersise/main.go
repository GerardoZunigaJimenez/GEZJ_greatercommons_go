package main

import (
	"fmt"
)

func main() {

	x:= []int {9,8,7,6,5,4,3,2,1,12,13,4,6,79,8,5,2}
	defer fmt.Printf("\nThe sum with variadic parameters of \n\t%v\n is\n\t%v ", x, sum(x...))

	fmt.Printf("\nThe sum with Slice of \n\t%v\n is\n\t%v",x, sumSlice(x))
}
func sum(s ...int)  int{
	total := 0
	for _, v := range s{
		total += v
	}

	return  total
}

func sumSlice(s []int)  int{
	total := 0
	for _, v := range s{
		total += v
	}

	return  total
}
