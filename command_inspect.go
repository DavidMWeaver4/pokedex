package main

import (
	"fmt"
)
func inspectCommand(cfg *config, pokemons []string) error{
	if len(pokemons) == 0{
		fmt.Println("no pokemon name given")
		return nil
	}
	name := pokemons[0]
	_, ok := cfg.Caught[name]
	if !ok{
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", cfg.Caught[name].Height)
	fmt.Printf("Weight: %d\n", cfg.Caught[name].Weight)
	fmt.Printf("Stats:\n")
	for _, sta:= range cfg.Caught[name].Stats{
		fmt.Printf("  - %s: %d\n", sta.Stat.Name, sta.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ:= range cfg.Caught[name].Types{
		fmt.Printf("  - %s\n", typ.Type.Name)
	}
	return nil
}
