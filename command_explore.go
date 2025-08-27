package main

import "fmt"

func commandExplore(cfg *config, params *[]string) error {
	if len(*params) < 1 {
		return fmt.Errorf("location name or id needed")
	}
	locationName := (*params)[0]
	location, err := cfg.pokeapiClient.GetLocation(&locationName)
	if err != nil {
		return err
	}
	//fmt.Print(PokemonList)
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
