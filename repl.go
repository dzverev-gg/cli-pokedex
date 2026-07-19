package main

import (
	"fmt"
	"os"
	"strings"
)
func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	result := strings.Fields(loweredText)

	return result
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

var registry map[string]cliCommand 

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage: \n\n")	
	for name, details := range registry {
		fmt.Printf("%s: %s\n", name, details.description)
	}
	return nil
}

func init() {

  registry = map[string]cliCommand {
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
}
}
