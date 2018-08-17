package main

import (
	"fmt"
	"sort"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortByAge []person
type sortByLast []person

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a sortByLast) Len() int           { return len(a) }
func (a sortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	fmt.Println(printUser( users...) )

	for _,v:= range users  {
		sort.Strings(v.Sayings)
	}

	fmt.Printf("\n\n\t*************** Sorted Users by %T ***************\n", sortByAge{})
	sort.Sort(sortByAge(users))
	fmt.Println(printUser( users...) )

	fmt.Printf("\n\n\t*************** Sorted Users by %T ***************\n", sortByLast{})
	sort.Sort(sortByLast(users))
	fmt.Println(printUser( users...) )
}

func printUser(people ...person) string {
	var message string
	for _, v := range people {
		message += fmt.Sprintln("User: ", v.First, v.Last, " with ", v.Age, " years old, hast the sayings: ")
		for _, s := range v.Sayings {
			message += fmt.Sprintln("\t\t", s)
		}
		message += "\n"
	}

	return message
}
