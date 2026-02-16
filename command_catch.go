package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func catchCommand(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no name provided.")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	poke, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	rolled := rand.IntN(poke.BaseExperience)
	if rolled > 40{
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.Caught[name] = poke
	}
	return nil
}
