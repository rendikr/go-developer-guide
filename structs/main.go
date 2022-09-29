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

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
