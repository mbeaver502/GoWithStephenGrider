package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// Create a person variable named alex.
	// Since we have not assigned initial values, the person fields are initialized to their zero values (empt string).
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v \n", alex)
}
