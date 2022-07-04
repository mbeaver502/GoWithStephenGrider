package main

// Note the syntax for importing multiple packages.
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of "deck", which is a slice of strings.
type deck []string

// Return a new deck of cards.
func newDeck() deck {
	cards := deck{} // Create a new empty deck.

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// _ is the discard variable.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// NB: This has no bounds-checking or any other kind of error handling...
func dealHand(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

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

func (d deck) toString() string {
	// Type-cast our deck to a slice of strings (i.e., a slice of slices of string).
	// Join together all the slices, separated by a comma.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read a deck from a file.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// Error handling for the result from ioutil.ReadFile
	if err != nil {
		// Option 1 -- Log the error and return a new deck (call newDeck())
		// Option 2 -- Log the error and entirely quit the program (fatal error)
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// This type conversion of deck(...) only works because deck is defined as a slice of string.
	// Note that strings.Split(...) returns a slice of string.
	return deck(strings.Split(string(bs), ","))
}
