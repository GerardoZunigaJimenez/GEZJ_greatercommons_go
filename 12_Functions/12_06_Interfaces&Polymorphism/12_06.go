package main

import "fmt"

//In Go, values can be of more than one type. An interface allows a value to be of more than one type. We create an interface
// using this syntax: “keyword identifier type” so for an interface it would be: “type human interface” We then
// define which method(s) a type must have to implement that interface. If a TYPE has the required methods, which could be
// none (the empty interface denoted by interface{}), then that TYPE implicitly implements the interface and is also
// of that interface type. In Go, values can be of more than one type.

// reference: https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secrete agent speak")
}

func (p person) speak() {
	fmt.Println(" I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	// a value can be more than one type
	sa1 := secretAgent{
		person{"James", "Bond"},
		true,
	}

	sa2 := secretAgent{
		person{"Miss", "Moneypenny"},
		true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	sa1.speak()
	sa2.speak()

	fmt.Println()

	bar(sa1)
	bar(sa2)
	bar(p1)

	fmt.Println()

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
