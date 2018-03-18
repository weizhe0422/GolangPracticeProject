package main

import (
	"fmt"
)

type contactInfo struct {
	email       string
	phoneNumber int64
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	ray := person{
		firstName: "WeiZhe",
		lastName:  "Chang",
		contactInfo: contactInfo{
			email:       "weizhe.chang@gmail.com",
			phoneNumber: 889200,
		},
	}

	ray.print()
	ray.updateFirstName("Ray")
	ray.print()
}

func (pointerPerson *person) updateFirstName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
