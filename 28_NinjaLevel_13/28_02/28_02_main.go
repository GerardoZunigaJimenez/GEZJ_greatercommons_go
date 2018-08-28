package main

import (
	"fmt"
	"GEZJ_greatercommons_go/28_NinjaLevel_13/28_02/quote"
	"GEZJ_greatercommons_go/28_NinjaLevel_13/28_02/word"
)

func main() {
	fmt.Println("String to be parsed:\n\t",quote.SunAlso)

	fmt.Println("\nWords Quantity: ", word.CountWordsByString(quote.SunAlso))

	fmt.Println("\nQuantity of repeated Words")
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println("\t",v, k)
	}
}
