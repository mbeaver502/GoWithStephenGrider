package main

import "fmt"

func main() {
	// Declare and assign a value to a variable called "card".
	//      var -- indicates we're declaring a variable
	//      card -- name of the variable
	//      string -- strong type of this variable, in this case a string
	//
	// Go is a statically typed language, like C++ or C#. This means we must specify the type.
	// In this case, the card variable will *always* contain string values.
	// Go has several basic types that we're familiar with: bool, string, int, float64, etc.
	//
	// We can be explicit:
	//      var card string = "Ace of Spades"
	//
	// Alternatively, we can let the Go compiler infer the strong type is a string:
	// 		card := "Ace of Spades"   	:= is only for assigning a value to a newly created variable!
	// 		card = "Five of Diamonds" 	Notice variable assignment does not use := for extant variables!

	card := newCard() // Even now the Go compiler can infer that newCard returns string, so card is string.

	fmt.Println(card)
}

// The newCard function has a return type of string. This function *must* return a string.
func newCard() string {
	return "Five of Diamonds"
}
