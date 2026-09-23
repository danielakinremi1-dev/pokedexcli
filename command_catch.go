package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Need a pokemon target")
	}

	targetPokemon := args[0]

	if len(targetPokemon) == 0 {
		return fmt.Errorf("No pokemon provided")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", targetPokemon)

	pokemonResp, err := cfg.pokeapiClient.PokemonCatch(targetPokemon)
	if err != nil {
		return err
	}

	catchOutcome := rand.Intn(pokemonResp.BaseExperience)

	if catchOutcome > 40 {
		fmt.Println(targetPokemon, "escaped!")
		return nil
	}

	fmt.Println(targetPokemon, "was caught!")
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.pokedex[pokemonResp.Name] = pokemonResp

	return nil
}
