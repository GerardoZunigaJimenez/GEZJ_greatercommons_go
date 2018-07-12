package main

import "fmt"

func main() {
	m := map[string] int{"James":32,"Miss Moneypenny":27,}

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

	fmt.Println()
	m["todd"] = 33
	for k, v := range m{
		fmt.Printf("Key: %v  value: %v\n",k,v)
	}

	fmt.Println()
	xi := []int{4,5,7,8,9,42}
	for i, v := range xi{
		fmt.Println(i,v)
	}
}
