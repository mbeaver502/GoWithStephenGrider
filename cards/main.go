package main

import "fmt"

// Main entrypoint for our program.
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
	//      card := "Ace of Spades"     := is only for assigning a value to a newly created variable!
	//      card = "Five of Diamonds"   Notice variable assignment does not use := for extant variables!
	//      card := newCard()           Even now the Go compiler can infer that newCard returns string, so card is string.
	//
	// Go has two data structures for representing lists of things:
	//      (1) Array -- a fixed-size list of things
	//      (2) Slice -- an array that can grow or shrink in size
	// Each element in an Array or Slice must be of the same type!
	//
	// Create a slice of strings representing our cards:
	// cards := []string{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades") // append returns a new slice object!
	//
	// The throwaway variable i represents the index of an item in the list.
	// The throwaway variable card is the current item in our list.
	// Both the i and card variables have local scope to this for-loop.
	// The range keyword allows us to iterate over the entire length of the cards slice.
	// Note: We *must* use the i variable somewhere since every declared variable must be used somewhere!
	//
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//

	cards := newDeck()
	// cards.print()

	hand, cards := dealHand(cards, 5)

	fmt.Println("-- Dealt Hand --")
	hand.print()
	fmt.Println("-- New Deck --")
	cards.print()
}

// The newCard function has a return type of string. This function *must* return a string.
func newCard() string {
	return "Five of Diamonds"
}
