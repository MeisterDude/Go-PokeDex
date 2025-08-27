package main

import "fmt"

func commandInspect(cfg *config, params *[]string) error {
	if len(*params) < 1 {
		return fmt.Errorf("pokemon name or id needed")
	}
	pokemonName := (*params)[0]
	p, ok := cfg.pokeapiClient.Pokedex.GetPokemon(pokemonName)
	if !ok {
		return fmt.Errorf("pokemon not caught or unknown")
	}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
