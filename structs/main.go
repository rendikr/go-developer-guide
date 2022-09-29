package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jimmy := person{
		firstName: "Jim",
		lastName:  "Morrison",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jimmy)
}
