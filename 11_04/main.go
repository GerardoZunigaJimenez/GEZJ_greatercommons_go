package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"MoneyPenny": 24,
			"Dr No":      45,
			"Q":          70,
		},
		favDrinks: []string{"martini", "rum", "coke",},
	}

	fmt.Println("Name: ",s.first)
	fmt.Println("Friends: ")
	for key, value := range s.friends {
		fmt.Printf("\tName: %v,\tAge:%v\n", key, value)
	}
	fmt.Println("Favorite Drinks: ")
	for _, value := range s.favDrinks {
		fmt.Printf("\t%v\n", value)
	}

}
