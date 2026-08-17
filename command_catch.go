package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	fmt.Println("Throwing a Pokeball at " + name + "...")
	if rand.Intn(100) > 75 {
		fmt.Println(name + " escaped!")
		return nil
	}

	pokemonRes, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Println(pokemonRes.Name + " got caught!")
	return nil
}