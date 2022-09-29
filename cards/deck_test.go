package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Error("Expected deck length of 20, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got ", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Error("Expected last card of Five of Clubs, but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Error("Expected deck length of 20, but got ", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got ", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "Five of Clubs" {
		t.Error("Expected last card of Five of Clubs, but got ", loadedDeck[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")
}
