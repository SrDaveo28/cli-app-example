package main

import (
	"log"
	"strconv"

	"github.com/DavidGenZ/cli-app-example/commands"
)

func main() {

	var expenses []float32
	for {
		input, err := commands.GetInput()

		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}

		espense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			continue
		}

		expenses = append(expenses, float32(espense))
	}
	commands.ShowInConsole(expenses)
}
