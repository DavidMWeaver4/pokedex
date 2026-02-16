package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var poke Pokemon
	err = json.Unmarshal(body, &poke)
	if err != nil {
		return Pokemon{}, err
	}
	return poke, nil
}
