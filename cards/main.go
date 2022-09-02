package main

func main() {
	// declare a new slice of type string
	// slice is a type of array that can grow or shrink, must have the same type of value
	// array is fixed length list
	cards := deck{
		newCard(),
		"Five of Diamonds",
	}

	// append a new value to slice
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
