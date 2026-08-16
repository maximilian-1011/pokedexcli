package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (locationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationArea{}, err
	}

	if res.StatusCode > 299 {
		return locationArea{}, fmt.Errorf("Http request failed with statuscode: %v", res.StatusCode)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	var location locationArea
	if err := json.Unmarshal(data, &location); err != nil {
		return locationArea{}, err
	}

	return location, nil
}
