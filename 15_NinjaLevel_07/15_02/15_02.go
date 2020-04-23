package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.age = 31
}

func (p *person) changeMe2() {
	p.first = "Sergio"
	p.last = "Jimenez"
	p.age = 24
}

func main() {
	p := person{
		first: "Gerardo",
		last:  "Zuniga",
		age:   24,
	}

	changeMe(&p)
	fmt.Println("p after 'func changeMe(p *person)':\t\t\t", p)
	p.changeMe2()
	fmt.Println("p after 'func (p *person) changeMe()':\t\t", p)
}
