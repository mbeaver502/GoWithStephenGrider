package main

import "fmt"

// Create a new type of "deck", which is a slice of strings.
type deck []string

// d is a receiver of type deck on the function print
// "d" is the alias we're using to refer to the actual deck variable that's using this function.
// Any variable that is of type deck has access to this print function.
// This would be similar to a public method on a class object, where all objects of the same type have access to that same public method.
// In a way, the "d" alias is similar to "this" or "self", but do *not* use those names.
// By convention, the receiver is named one or two characters based on the type (so "deck" => "d").
// Note that Go is NOT object-oriented, so this is not a method on a class!
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
