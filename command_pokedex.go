package main

import (
	"fmt"
)

func pokedexCommand(cfg *config, name []string)error{
	if len(cfg.Caught) == 0{
		fmt.Println("Your Pokedex is empty.")
		return nil
	}
	fmt.Println("Your Pokedex: ")
	for _, pokes := range cfg.Caught{
		fmt.Printf(" - %s\n", pokes.Name)
	}
	return nil
}
