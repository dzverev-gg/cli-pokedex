package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputString := scanner.Text()
		command, exists := registry[inputString]
		if exists {
			command.callback()
		} else {
			fmt.Print("Unknown command \n")
		}

	}
}

