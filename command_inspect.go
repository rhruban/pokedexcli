package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single Pokemon name")
	}

	name := args[0]
	poks, exists := cfg.caughtPokemon[name]
	if !exists {
		return errors.New("you have not caught this Pokemon")
	}

	fmt.Printf("Name: %s\n", poks.Name)
	fmt.Printf("Height: %d\n", poks.Height)
	fmt.Printf("Weight: %d\n", poks.Weight)
	fmt.Printf("Stats:\n")

	for _, stat := range poks.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range poks.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}
	return nil
}
// Name: pidgey
// Height: 3
// Weight: 18
// Stats:
//   -hp: 40
//   -attack: 45
//   -defense: 40
//   -special-attack: 35
//   -special-defense: 35
//   -speed: 56
// Types:
//   - normal
//   - flying
