package main

import (
	"fmt"
	"github.com/maximilian-1011/pokedexcli/internal/pokeapi"
)

func commandMapb(state *config) error {
	if state.previous == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	location, err := pokeapi.GetLocationArea(*state.previous)
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
