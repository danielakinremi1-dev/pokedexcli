package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Need a location argument")
	}

	location := args[0]

	if len(location) == 0 {
		return fmt.Errorf("No location provided")
	}

	locationResp, err := cfg.pokeapiClient.LocationDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
