package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First   string   `json:"Name"`
	Last    string   `json:"LastName"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	var people []user
	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()
	for _, v := range people {
		fmt.Println("User: ", v.First, v.Last, " with ", v.Age, " years old, hast the sayings: ")
		for _, s := range v.Sayings {
			fmt.Println("\t\t", s)
		}
		fmt.Println()
	}
}
