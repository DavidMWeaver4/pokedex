package main

import (
	"fmt"

)

func exploreCommand(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location provided")
	}
	name := args[0]
	fmt.Println("Exploring " + name + "...")
	locationEncounters, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, le := range locationEncounters.PokemonEncounters {
		fmt.Println(" - " + le.Pokemon.Name)
	}
	return nil
}
