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
	}

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// send the function to go routines
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// assign a value to channel
		c <- "might be down I think"
		return
	}

	fmt.Println(link, "is up and running!")
	// assign a value to channel
	c <- "yep it's up"
}
