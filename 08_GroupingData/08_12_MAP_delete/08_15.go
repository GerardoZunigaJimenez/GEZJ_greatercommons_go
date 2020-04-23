package main

import "fmt"

func main() {
	m := map[string]int{"James": 32, "Miss Moneypenny": 27}

	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	fmt.Println()

	delete(m, "I'm Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Flemming"])

	if _, ok := m["Ian Flemming"]; ok {
		fmt.Println(ok)
	}
}
