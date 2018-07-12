package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	lenght float64
	with   float64
}

type circle struct {
	radio float64
}

func (r rectangle) area() float64 {
	return r.lenght * r.with
}

func (c circle) area() float64 {
	return (c.radio * c.radio) * math.Pi
}

type shape interface{
	area() float64
}

func info(s shape) string{
	return fmt.Sprintf("The shape %T has the area: %v\n", s, s.area() )
}

func main() {
	r := rectangle{
		lenght:4,
		with:3,
	}

	c := circle{
		radio: 5,
	}

	fmt.Println( info(r) )
	fmt.Println( info(c) )
}
