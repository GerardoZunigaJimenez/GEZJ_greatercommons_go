package main

import (
	"fmt"
	"time"
)

type customErr struct {
	s string
	d time.Time
	l time.Location
}

func (ce customErr) Error() string {
	return fmt.Sprintf( "there was an error %v at the time: %v with the location %v", ce.s, ce.d, ce.l)
}

func main() {
	c1 := customErr{
		s:"marshal error at the name",
		d:time.Now(),
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
