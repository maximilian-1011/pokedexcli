package main

import (
	"fmt"
	"github.com/maximilian-1011/pokedexcli/internal/pokeapi"
)
func commandMap(state *config) error {
	location, err := pokeapi.GetLocationArea(state.next)
	if err != nil {
		return err
	}

	for i := 0; i < len(location.Results); i++ {
		fmt.Println(location.Results[i].Name)
	}

	state.previous = location.Previous
	state.next = location.Next
	return nil
}
