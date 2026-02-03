package main

import(
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locs, err := cfg.pokeapiClient.ListLocations(cfg.nextLocURL)
	if err != nil {
		return err
	}

	cfg.prevLocURL = locs.Previous
	cfg.nextLocURL = locs.Next

	for i := range locs.Results {
		fmt.Println(locs.Results[i].Name)
	}

	return nil	
}

func commandMapb(cfg *config) error {
	if cfg.prevLocURL == nil {
		return errors.New("you're on the first page")
	}

	locs, err := cfg.pokeapiClient.ListLocations(cfg.prevLocURL)
	if err != nil {
		return err
	}

	cfg.prevLocURL = locs.Previous
	cfg.nextLocURL = locs.Next

	for i := range locs.Results {
		fmt.Println(locs.Results[i].Name)
	}

	return nil	
}
