package main

import(
	"fmt"
	"github.com/rhruban/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.nextLocURL == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	locs := pokeapi.GetMap(cfg.nextLocURL)
	for i := range locs.Results {
		fmt.Println(locs.Results[i].Name)
	}
	cfg.prevLocURL = locs.Previous
	cfg.nextLocURL = locs.Next
	return nil	
}

func commandMapb(cfg *config) error {
	if cfg.prevLocURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locs := pokeapi.GetMap(cfg.prevLocURL)
	for i := range locs.Results {
		fmt.Println(locs.Results[i].Name)
	}
	cfg.prevLocURL = locs.Previous
	cfg.nextLocURL = locs.Next
	return nil	
	fmt.Println("Did a mapb")
	return nil	
}
