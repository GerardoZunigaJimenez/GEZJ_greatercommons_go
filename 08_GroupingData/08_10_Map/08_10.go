package main

import "fmt"

func main() {
	m := map[string]int{"James": 32, "Miss Moneypenny": 27}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	fmt.Println()

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println()

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("this is the if print", v)
	}

	fmt.Println("SEarching: ", m["geras"])
}
