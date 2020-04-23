package main

import (
	"log"
	"os"
)

/* Panic is equivalent to println followed by a call to panic()
 */
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
}
