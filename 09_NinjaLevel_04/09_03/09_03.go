package main

import "fmt"

func main() {
	a := []int{42,43,44,45,46,47,48,49,50,51}

	s1 := a[0:5]
	s2 := a[5:10]
	s3 := a[2:7]
	s4 := a[1:6]


	for p,i:= range s1{
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n",p,i,i)
	}
fmt.Println()
	for p,i:= range s2{
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n",p,i,i)
	}
	fmt.Println()
	for p,i:= range s3{
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n",p,i,i)
	}
	fmt.Println()
	for p,i:= range s4{
		fmt.Printf("The position [%v] hast a type '%T' witht the value: %v\n",p,i,i)
	}
}
