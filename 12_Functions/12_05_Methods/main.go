package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func ( r receiver) identifier (parameters) return(s) { }
func main() {
	sa1 := secretAgent{
		person{"James", "Bond"},
		true,
	}

	sa2 := secretAgent{
		person{"Miss", "Moneypenny"},
		true,
	}

	sa1.speak()
	sa2.speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
