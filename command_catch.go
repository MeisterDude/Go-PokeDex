package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, params *[]string) error {
	if len(*params) < 1 {
		return fmt.Errorf("pokemon name or id needed")
	}
	pokemonName := (*params)[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(&pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	chance := rand.IntN(700)
	if chance >= pokemon.BaseExperience {
		cfg.pokeapiClient.Pokedex.AddPokemon(pokemonName, pokemon)
		fmt.Printf("Successfully caught %s...\n", pokemonName)
	} else {
		fmt.Printf("%s escaped...\n", pokemonName)
	}
	return nil
}
