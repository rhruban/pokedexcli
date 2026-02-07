package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/rhruban/pokedexcli/internal/pokeapi"
)

func enterREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		cleanInputs := cleanInput(userInput)
		if len(cleanInputs) == 0 {
			continue
		}

		commandName := cleanInputs[0]
		args := []string{}
		if len(cleanInputs) > 1 {
			args = cleanInputs[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unkown command")
			continue
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type config struct {
	pokeapiClient pokeapi.Client
	prevLocURL *string
	nextLocURL *string
}

type cliCommand struct {
	name				string
	description	string
	callback		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays next 20 location areas",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 location areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Display pokemon that can be encountered in a location area",
			callback: commandExplore,
		},
	}
}

