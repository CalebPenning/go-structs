package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "P",
		contact: contactInfo{
			email: "Jim@gmoil.corm",
			zipCode: 94550,
		},
	}
}