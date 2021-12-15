package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// 
	jim := person {
		firstName: "Jim",
		lastName: "P",
		contactInfo: contactInfo{
			email: "Jim@gmoil.corm",
			zipCode: 94550,
		},
	}

	fmt.Println("I'm am Jim still")
	jim.print()
	// &variable == give me the memory address of this value
	jimPointer := &jim
	jimPointer.updateFirstName("JAMES>LOL")
	jim.print()
}

// *pointer == give me the value that is sitting in that memory address
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// this line doesnt work? p.firstName = newFirstName
	// we must utilize pointers

	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}