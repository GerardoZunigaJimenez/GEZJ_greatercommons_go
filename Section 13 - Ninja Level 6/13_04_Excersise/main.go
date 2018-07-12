package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf("My name is %v %v and my age is %v", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Gerardo",
		last:  "Zuniga",
		age:   31,
	}.speak()
	fmt.Println(p1)

}
