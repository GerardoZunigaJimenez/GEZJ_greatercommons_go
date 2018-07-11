package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		"James",
		"Bond",
		[]string{"vanilla", "chocolate"},
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		[]string{"strawberry", "peach"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range m {
		fmt.Println(p.first, p.last)
		for _, v := range p.iceCream {
			fmt.Printf("\t%v", v)
		}
		fmt.Println("\n")
	}
}
