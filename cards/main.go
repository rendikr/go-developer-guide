package main

import "fmt"

func main() {
	// declare a variable

	// this one works
	// var card string = "Ace of Spades"

	// this one also works
	card := newCard()

	// reassign a new value to a variable
	// card = "Five of Diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
