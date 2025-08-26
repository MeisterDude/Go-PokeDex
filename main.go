package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

func main() {
	repl()
}
