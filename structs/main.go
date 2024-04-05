package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "sergiomondragondev@gmail.com",
			zipCode: 12345,
		},
	}
	// var alex person
	// alexPointer := &alex
	//  REFERENCE TO THE MEMORY ADDRESS
	// alexPointer.updateLastName("Mondragon")
	// alexPointer.updateName("Sergio")
	alex.updateName("Sergio")
	alex.updateLastName("Mondragon")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// Value of the memory addres is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("First name %s and last name %s and contact info: %+v\n", p.firstName, p.lastName, p.contact)
}
