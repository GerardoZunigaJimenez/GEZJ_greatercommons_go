package main

import "fmt"

type person struct {
	first   string
	message string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.first, "says '", p.message, "'")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first:   "Gerardo",
		message: "Go to Climb",
	}

	saySomething(&p)
}
