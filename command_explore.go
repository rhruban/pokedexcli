package main

import(
	"fmt"
)

func commandExplore(cfg *config, arg1 *string) error {
	locs, err := cfg.pokeapiClient.ListLocationPokemon(arg1)
	if err != nil {
		return err
	}

	for i := range locs.PokemonEncounters {
		fmt.Println(locs.PokemonEncounters[i].Pokemon.Name)
	}

	return nil	
}
