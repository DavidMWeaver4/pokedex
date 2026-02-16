package pokeapi
import (
	"encoding/json"
	"io"
	"fmt"
)

func (c *Client) GetLocationArea(locationAreaName string)(LocationEncounters, error){
	url := baseURL + "/location-area/" + locationAreaName
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationEncounters{},err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Errorf("bad status code: %d", res.StatusCode)
		return LocationEncounters{}, err
	}
	if err != nil {
		return LocationEncounters{},err
	}
	var locationEncount LocationEncounters
	err = json.Unmarshal(body, &locationEncount)
	if err != nil {
		return LocationEncounters{},err
	}
	return locationEncount, nil
}
