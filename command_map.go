package main

import (
	"fmt"

	"github.com/DavidMWeaver4/pokedex/internal/pokeapi"
)

func mapCommand(cfg *config, args []string) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	}
	res, err := pokeapi.ListLocationAreas(url)
	if err != nil {
		return err
	}
	if res.Next != nil {
		cfg.Next = res.Next
	} else {
		cfg.Next = nil
	}
	if res.Previous != nil {
		cfg.Prev = res.Previous
	} else {
		cfg.Prev = nil
	}
	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCommand(cfg *config, args []string) error {
	url := ""
	if cfg.Prev == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.Prev
	}
	res, err := pokeapi.ListLocationAreas(url)
	if err != nil {
		return err
	}
	if res.Next != nil {
		cfg.Next = res.Next
	} else {
		cfg.Next = nil
	}
	if res.Previous != nil {
		cfg.Prev = res.Previous
	} else {
		cfg.Prev = nil
	}
	for _, result := range res.Results {
		fmt.Println(result.Name)
	}
	return nil
}
