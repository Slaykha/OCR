package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("images.png")
	text, err := client.Text()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(text)

	// Hello, World!
}
