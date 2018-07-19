package main

import "fmt"

func main() {
	favSport := "climbing"
	switch favSport{
	case "soccer":
		fmt.Println("soccer")
	case "basquetball":
		fmt.Println("basquetball")
	case "hockey":
		fmt.Println("hockey")
	case "climbing":
	fmt.Println("climbing")
	default:
		fmt.Println("climbing is the only real sport")
	}

}
