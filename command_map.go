package main

import(
	"fmt"
	"github.com/rhruban/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	locs := pokeapi.GetMap(url)
	fmt.Printf("\n%v\n",locs)
	fmt.Println("Did a map")
	return nil	
}

func commandMapb() error {
	fmt.Println("Did a mapb")
	return nil	
}
