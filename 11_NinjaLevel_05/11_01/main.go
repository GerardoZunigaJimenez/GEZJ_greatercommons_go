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
		[]string {"vanilla", "chocolate"},
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		[]string {"strawberry", "peach"},
	}

	fmt.Println(p1.first, p1.last)
	for _,v := range p1.iceCream {
		fmt.Printf("\t%v",v)
	}

	fmt.Println()

	fmt.Println(p2.first, p2.last)
	for _,v := range p2.iceCream {
		fmt.Printf("\t%v",v)
	}
}
