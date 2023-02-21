package main

import (
	"fmt"
	"log"

	"example.com/m/v2/commands"
)

func main() {

	for {
		input, err := commands.GetInput()

		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}
	}
	fmt.Println("Closing ...")
}
