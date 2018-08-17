package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type personWithErrors struct{
	person
	err error
}

func (p personWithErrors) Error() string {
	return fmt.Sprintf("The user %v %v had problems to be translated at: %v", p.First, p.Last, p.err)
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln( personWithErrors{ p1, err})
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("There was an error marshalling %v in toJSON: %v", a, err)
	}
	err = errors.New( fmt.Sprintf("Provoked Failure at attribute %v", a))
	return bs, err
}
