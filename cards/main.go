package main

import "fmt"

func main() {
	// declare a new slice of type string
	// slice is a type of array that can grow or shrink, must have the same type of value
	// array is fixed length list
	cards := []string{
		newCard(),
		"Five of Diamonds",
	}

	// append a new value to slice
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
