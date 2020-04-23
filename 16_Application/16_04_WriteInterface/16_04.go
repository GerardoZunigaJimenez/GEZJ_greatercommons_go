package main

import "fmt"
import (
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Word")
	fmt.Fprintln(os.Stdout, "Hello World")
	io.WriteString(os.Stdout, "Hello World")
}
