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

	// jimPointer := &jim, way to get memory address
	// however, passing in jim, while taking *person rather than just person
	// go takes a shortcut thru the "pass by value" nature of itself
	// and just updates the value AT the memory address of that struct
	jim.updateFirstName("JAMES>LOL")
	jim.print()
}

// *pointer == give me the value that is sitting in that memory address
// this is a type description; it means we're working with a pointer to a type
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// this line doesnt work? p.firstName = newFirstName
	// we must utilize pointers

	// this vvv is an operator, it means we want to change the value the pointer is 'pointing' to
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}


// turn address into value with *address
// turn value into address with &value