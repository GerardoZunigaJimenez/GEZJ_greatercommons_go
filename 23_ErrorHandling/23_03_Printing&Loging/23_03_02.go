package main

import (
	"os"
	"fmt"
	"log"
)

//Print calls Output to print to the standard logger.  Arguments are handle in the manner of fmt.Println
func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)

	f2, err:= os.Open("no-file.txt")
	if err != nil {
		log.Println("log err happened", err)
		panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
