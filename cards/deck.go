package main

import "fmt"

// Create a new type of "deck", which is a slice of strings.
type deck []string

// d is a receiver of type deck on the function print
// Any variable that is of type deck has access to this print function.
// This would be similar to a public method on a class object, where all objects of the same type
//  have access to that same public method.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
