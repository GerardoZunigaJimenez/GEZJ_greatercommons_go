package main

import "fmt"

func main() {

	m := map[string] []string{}

	m[`bond_james`] = []string {`Shaken, not stirred`, `Martinis`, `Women`}
	m[`moneypenny_miss`] = []string {`James Bond`, `Literature`, `Computer Science`}
	m[`no_dr`] = []string {`Being evil`, `Ice cream`, `Sunsets`}

	for i, v1 := range m{
		fmt.Println("record: ", i)
		for j, v2 := range v1 {
			fmt.Printf("\t index position: %v \t value: \t %v\n",j,v2)
		}
	}
}
