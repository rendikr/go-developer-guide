package main

func main() {
	// declare a new slice of type string
	// slice is a type of array that can grow or shrink, must have the same type of value
	// array is fixed length list
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
