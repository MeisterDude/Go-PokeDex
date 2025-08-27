package main

import "fmt"

func commandPokedex(cfg *config, params *[]string) error {
	if cfg.pokeapiClient.Pokedex.IsEmpty() {
		return fmt.Errorf("no pokemon caught")
	}
	fmt.Printf("Your Pokedex:\n")
	for _, v := range cfg.pokeapiClient.Pokedex.GetPokedex() {
		fmt.Printf("  -  %s\n", v.Name)
	}
	return nil
}
