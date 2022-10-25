package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jimmy := person{
		firstName: "Jimmy",
		lastName:  "Morrison",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 94000,
		},
	}

	jimmy.print()
	jimmy.updateName("Jim")
	jimmy.print()
}

// *person = type pointer to a person
// *pointerToPerson = operator to manipulate the value this pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
