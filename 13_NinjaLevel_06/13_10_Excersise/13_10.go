package main

import "fmt"

func main() {
	m := map[string] int {}
	m["a"] = 5
	m["b"] = 8
	m["d"] = 3
	m["c"] = 16

	for k,v := range m {
		newV := 0
		f := incrementor()
		for i := 0; i < v ; i++ {
			a := f()
			newV = a
		}
		m[k] = newV
	}

	fmt.Println()
	for k,v := range m {
		fmt.Printf("The execution '%v' have been executed '%v' times\n", k, v)
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}