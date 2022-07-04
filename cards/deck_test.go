// The xxx_test.go filename tells us that this code exists specifically to test the xxx code.
// We can run `go test` to run tests.
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be an Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be a King of Diamonds, but got %v", d[len(d)-1])
	}
}
