package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	num := len(cfg.caughtPokemon)
	if num == 0 {
		return errors.New("you have not caught any Pokemon")
	}
	fmt.Printf("Your Pokedex:\n")
	for key := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n",key)
	}
	return nil
}
