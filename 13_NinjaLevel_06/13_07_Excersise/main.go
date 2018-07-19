package main

import "fmt"

func main() {
	f := func(s string) string {
		return fmt.Sprintf("This is the message that we are creating: '%v'\n", s)
	}

	fmt.Println( f("I'm Gerardo Zuniga ") )
	fmt.Printf( "f type: '%T'\n", f)
}