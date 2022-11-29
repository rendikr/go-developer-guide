package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type numbers []int

func newNumbers() numbers {
	listOfNumbers := numbers{}

	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numberList {
		listOfNumbers = append(listOfNumbers, num)
	}

	return listOfNumbers
}

// allow any value with the type of "deck" to call this function. this is called "Receiver"
func (n numbers) print() {
	for _, num := range n {
		if (num %2 == 0) {
			fmt.Println(num, "is an Even number")
		} else {
			fmt.Println(num, "is an Odd number")
		}
	}
}
