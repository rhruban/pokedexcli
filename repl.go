package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func enterREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := getConfig()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInputs := cleanInput(userInput)
		if len(cleanInputs) == 0 {
			continue
		}
		commandName := cleanInputs[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unkown command")
			continue
		} else {
			err := command.callback(cfg)
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
	prevLocURL *string
	nextLocURL *string
}

func getConfig() *config {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := config{
		prevLocURL: nil,
		nextLocURL: &baseURL,
	}
	return &cfg
}

type cliCommand struct {
	name				string
	description	string
	callback		func(*config) error
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
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 location areas",
			callback: commandMapb,
		},
	}
}

