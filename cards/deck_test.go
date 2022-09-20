package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Error("Expected deck length of 20, but got ", len(d))
	}
}
