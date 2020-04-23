package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type user struct {
	ID      int
	Name    string
	Last    string
	Mail    string
	Birth   time.Time
	Enabled bool
}

func main() {
	uList := []user{
		user{ID: 1, Name: "Gerardo", Last: "Zuniga", Mail: "gerardo.zuniga@hotmail.com", Birth: time.Date(1986, time.October, 03, 00, 00, 00, 00, time.UTC), Enabled: true},
		user{ID: 2, Name: "Sergio", Last: "Silva", Mail: "ssilva@hotmail.com", Birth: time.Date(1994, time.March, 04, 00, 00, 00, 00, time.UTC), Enabled: true},
		user{ID: 3, Name: "Tony", Last: "Ramos", Mail: "tony.ramos@go.net", Birth: time.Date(1986, time.November, 18, 00, 00, 00, 00, time.UTC), Enabled: true},
		user{ID: 4, Name: "Ana Eugenia", Last: "Rios", Mail: "ana.rios@ab&b.com", Birth: time.Date(1987, time.May, 24, 00, 00, 00, 00, time.UTC), Enabled: true},
		user{ID: 5, Name: "Tulio", Last: "Santander", Mail: "prime.cuevas.18@arquitectos.com", Birth: time.Date(1994, time.September, 16, 00, 00, 00, 00, time.UTC), Enabled: true},
	}

	fmt.Println(printUser(uList...))

	bs, err := json.Marshal(uList)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
}

func printUser(uList ...user) string {
	var message string
	for _, v := range uList {
		if v.Enabled {
			message += fmt.Sprintf("The user #%v called %v %v, with contact id '%v' borned at %v\n", v.ID, v.Name, v.Name, v.Mail, v.Birth.Local())
		}
	}
	return message
}
