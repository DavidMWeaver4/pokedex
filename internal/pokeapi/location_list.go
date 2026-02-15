package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocationAreas(pageurl string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageurl != "" {
		url = pageurl
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	if err != nil {
		return LocationAreasResponse{}, err
	}
	var locationAreasResp LocationAreasResponse
	err = json.Unmarshal(body, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	return locationAreasResp, nil
}
