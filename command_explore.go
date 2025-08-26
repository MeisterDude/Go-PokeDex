package main

import "fmt"

func commandExplore(cfg *config, params *[]string) error {
	if len(*params) < 1 {
		return fmt.Errorf("location name or id needed")
	}
	location := (*params)[0]
	PokemonList, err := cfg.pokeapiClient.ListPokemon(&location)
	if err != nil {
		return err
	}
	//fmt.Print(PokemonList)
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", location)
	for _, pokemon := range PokemonList.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
