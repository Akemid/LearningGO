package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove any existing test files
	os.Remove("_decktesting")

	// Create a new deck
	deck := newDeck()

	// Save the deck to a file
	deck.saveToFile("_decktesting")

	// Load the deck from the file
	loadedDeck := readFromFile("_decktesting")

	// Check if the deck has 52 cards
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	// Remove the test file
	os.Remove("_decktesting")
}
