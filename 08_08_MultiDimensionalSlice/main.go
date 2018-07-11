package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"miss","moneyPenny","Strawberry","Hazelnut"}
	fmt.Println(mp)

	fmt.Println()

	xp:= [][]string{jb,mp}
	fmt.Println(xp)

}
