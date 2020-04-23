package main

import "fmt"

func main() {

	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	s3 := [][]string{s1, s2}
	fmt.Println(s3)

	for i, v := range s3 {
		fmt.Println("record: ", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, v2)
		}
	}
}
