package main

import (
	"os"
	"log"
)

func main(){
	_, err := os.Open("no-file.txt")
	if err != nil{
		log.Fatal(err)
	}
}

//the fatal functions call os.Exit(1) after writing the log message
//fatalln is equivalent to Println() followed by a call to os.Exit(1)