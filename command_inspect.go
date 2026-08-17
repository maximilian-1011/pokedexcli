package main

import (
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	pokemonResp, ok := cfg.pokeapiClient.DexGet(name)
	if !ok {
		fmt.Println("The Pokemon has not been caught yet")
		return nil
	}

	fmt.Println("Name: "+pokemonResp.Name)
	fmt.Printf("Height: %v\n", pokemonResp.Height)
	fmt.Printf("Weight: %v\n", pokemonResp.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonResp.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	} 
	fmt.Println("Types:")
	for _, tp := range pokemonResp.Types {
		fmt.Printf(" - %v\n", tp.Type.Name)
	}

	return nil
}