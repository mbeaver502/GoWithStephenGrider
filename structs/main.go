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

	johnPointer := &john
	johnPointer.updateName("Jim")
	john.print()
}

// Go is a pass-by-value language.
// So when we say we want to use p, Go needs to copy the value somewhere in memory so that we can access it.
// This is unlike C#, where non-scalar variables are always passed by reference.
// So if we try to modify p inside the print() function, we are modifying only the copy, not the original.
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// If we want to modify p directly, we can pass it by reference by using a pointer.
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
