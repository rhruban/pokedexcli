package main

import (
	"time"
	"github.com/rhruban/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	enterREPL(cfg)
}
