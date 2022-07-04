package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// Create a person variable named alex.
	// Since we have not assigned initial values, the person fields are initialized to their zero values (empt string).
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "john@example.com"
	// alex.contact.zip = "20007"

	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email: "john@example.com",
			zip:   "12345",
		},
	}

	john.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
