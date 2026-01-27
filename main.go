package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		fmt.Printf("Your command was: %s\n", cleanInput[0])
	}
}
