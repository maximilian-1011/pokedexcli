package main

import (
	"fmt"
)

func commandPokedex(cfg *config, argument string) error {
	pokedex := cfg.pokeapiClient.GetDex()

	for _, pokemon := range pokedex {
		fmt.Println(" - " + pokemon.Name)
	}

	return nil
}