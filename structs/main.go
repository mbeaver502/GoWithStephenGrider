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

	// johnPointer := &john // & gives us the address of the john variable. This gives us a pointer to john.
	// johnPointer.updateName("Jim")
	(&john).updateName("Jimmy")
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
// (p *person() here tells us that this function can be used by any variable that is a pointer to a person type.
func (p *person) updateName(newFirstName string) {
	// By using *p, we can access the value stored at the address referenced by p.
	// If we used just p (instead of *p), then we'd be operating on the address itself, not the stored value.
	// (*p) can be interpreted as the value stored at the memory referenced by the pointer p.
	// So (*p) gets us the struct stored in memory at the address referenced by the pointer p.
	(*p).firstName = newFirstName
}
