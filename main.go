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
		cleanedText := cleanInput(inputString)
		fmt.Printf("Your command was: %s \n", cleanedText[0])
	}
}

