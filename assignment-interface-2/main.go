package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file", fileName)
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, file)
}
