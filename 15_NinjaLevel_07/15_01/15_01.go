package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"i": 32,
		"s": "Hello World",
		"f": 56.25,
		"b": true,
	}

	i := 32
	s := "Hello World"
	f := 56.25
	b := true

	fmt.Println("Printing Variables one by one")
	fmt.Printf("\tThe variable '%T' with the address '%v' has the value '%v'\n", i, &i, i)
	fmt.Printf("\tThe variable '%T' with the address '%v' has the value '%v'\n", s, &s, s)
	fmt.Printf("\tThe variable '%T' with the address '%v' has the value '%v'\n", f, &f, f)
	fmt.Printf("\tThe variable '%T' with the address '%v' has the value '%v'\n", b, &b, b)

	fmt.Println("\nPrinting variables with a map")
	for k, v := range m {
		fmt.Printf("\tThe variable '%v' with the type '%T' and the address '%v' has the value '%v'\n", k, v, &v, v)
	}
}
