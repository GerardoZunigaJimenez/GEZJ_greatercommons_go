package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5,7,89,24,51,3,9,8,10,12,}
	xs := []string{"James","Q","M","Moneypenny","Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-------------------------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)




}
