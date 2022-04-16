package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected K of Club, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {

}
