package main

import(
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single location name")
	}

	name := args[0]
	locs, err := cfg.pokeapiClient.ListLocationPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locs.Name)
	fmt.Println("Found Pokemon: ")
	for i := range locs.PokemonEncounters {
		fmt.Printf(" - %s\n", locs.PokemonEncounters[i].Pokemon.Name)
	}

	return nil	
}
