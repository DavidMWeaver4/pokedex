package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationEncounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func exploreCommand(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location provided")
	}
	name := args[0]
	fmt.Println("Exploring " + name + "...")
	fmt.Println("Found Pokemon:")
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + name)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", res.StatusCode)
	}
	if err != nil {
		return err
	}
	var locationEncounters LocationEncounters
	err = json.Unmarshal(body, &locationEncounters)
	if err != nil {
		return err
	}
	for _, le := range locationEncounters.PokemonEncounters {
		fmt.Println(" - " + le.Pokemon.Name)
	}
	return nil
}
