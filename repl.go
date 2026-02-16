package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DavidMWeaver4/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Prev          *string
	Caught        map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned_text := cleanInput(scanner.Text())
		if len(cleaned_text) == 0 {
			continue
		}
		command, exists := getCommands(cfg)[cleaned_text[0]]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, cleaned_text[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	temp_text := strings.ToLower(strings.TrimSpace(text))
	edit_text := strings.Fields(temp_text)
	return edit_text
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands(cfg *config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "Displays all pokemon in a location",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Attemps to catch a pokemon",
			callback:    catchCommand,
		},
		"inspect"{
			name:			"inspect",
			description:	"Inspect a pokemon you have caught",
			callback: 		inspectCommand,
		},
	}
}
