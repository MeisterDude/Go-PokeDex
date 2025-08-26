package main

import (
	"fmt"
)

func commandHelp(cfg *config, params *[]string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cc := range getCommands() {
		fmt.Printf("%s: %s\n", cc.name, cc.description)
	}
	return nil
}
