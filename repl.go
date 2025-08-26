package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MeisterDude/Go-PokeDex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type clientCommand struct {
	name        string
	description string
	callback    func(*config, *[]string) error
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
		"map": {
			name:        "map",
			description: "Get the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get the pokemon from a given location",
			callback:    commandExplore,
		},
	}
}

func repl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]
		params := input[1:]
		//fmt.Printf("Your command was: %s\n", cleanedInput[0])
		if cc, ok := getCommands()[command]; ok {
			if err := cc.callback(cfg, &params); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower((text))
	return strings.Fields(lowerText)
}
