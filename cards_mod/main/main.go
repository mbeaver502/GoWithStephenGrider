package main

import (
	"mbeaver/deck"
)

func main() {
	cards := deck.NewDeck()

	cards.Shuffle()
	cards.Print()
}
