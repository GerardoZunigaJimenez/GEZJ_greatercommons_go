package main

import (
	"os"
	"fmt"
	"log"
)

func main(){
	_, err := os.Open("no-file.txt")
	if err != nil{
		//fmt.Println("fmt: err happened", err)
		//log.Println("log err happened", err)
		log.Fatalln("fatal err happened",err)
		log.Panicln("panic err append", err)

		fmt.Println("I'm still alive")

		panic(err)
	}
}

/*
the fatal functions call os.Exit(1) after writing the log message
 */
 //Fatalln is equivalent to Println followed by a call to ox.Exit(1)