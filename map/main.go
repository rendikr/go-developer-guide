package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#00ff00",
	// }

	// assign a new key value to the colors map
	colors["white"] = "#ffffff"

	// delete an item from colors map by key
	delete(colors, "white")


	fmt.Println(colors)
}
