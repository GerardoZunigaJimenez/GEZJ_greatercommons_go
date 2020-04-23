package main

import (
	"fmt"
	"log"
	"os"
)

//the fatal functions call os.Exit(1) after writing the log message
//fatalln is equivalent to Println() followed by a call to os.Exit(1)
func main() {

	defer foo()
	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("when Os.exit(1) is called, defered functions don't run")
}
