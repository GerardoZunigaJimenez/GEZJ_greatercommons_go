package main

import (
	"log"
	"fmt"
)

/*
We can add information to our errors. We can do this with
	errors.New()
		fmt.Errorf()
	builtin.error
“Error values in Go aren’t special, they are just values like any other,
and so you have the entire language at your disposal.” - Rob Pike
 */
func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, errNorgateMath
	}
	return 42, nil
}
