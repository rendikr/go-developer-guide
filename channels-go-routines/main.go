package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://localhost:3000",
	}

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// send the function to go routines
		go checkLink(link, c)
	}

	// eternity for loop
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// assign a value to channel
		c <- link
		return
	}

	fmt.Println(link, "is up and running!")
	// assign a value to channel
	c <- link
}
