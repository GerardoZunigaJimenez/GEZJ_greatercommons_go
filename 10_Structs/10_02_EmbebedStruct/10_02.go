package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			32,
		},
		ltk: true,
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		28,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println()
	fmt.Println("with sub attributes definition: ", sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)
	fmt.Println("without sub attibutes definition: R", sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
