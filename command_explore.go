package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationName string) error {
	if locationName == "" {
		return fmt.Errorf("No location provided")
	}

	locationResp, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		return err
	}

	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}