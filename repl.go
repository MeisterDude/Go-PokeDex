package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type clientCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]clientCommand {
	return map[string]clientCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func repl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]
		//fmt.Printf("Your command was: %s\n", cleanedInput[0])
		if cc, ok := getCommands()[command]; ok {
			cc.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cc := range getCommands() {
		fmt.Printf("%s: %s\n", cc.name, cc.description)
	}
	return nil
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower((text))
	return strings.Fields(lowerText)
}
