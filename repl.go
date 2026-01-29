package main

import(
	"strings"
	"fmt"
	"os"
	"bufio"
)

type cliCommand struct {
	name				string
	description	string
	callback		func() error
}

var commands = map[string]cliCommand{}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}

func enterREPL() {
	commands = map[string]cliCommand{
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
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInputs := cleanInput(userInput)
		command, ok := commands[cleanInputs[0]]
		if !ok {
			fmt.Println("Unkown command")
		} else {
			command.callback()
			fmt.Printf("Your command was: %s\n", cleanInputs[0])
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
