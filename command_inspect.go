package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Need a pokemon to inspect")
	}

	pokemon := args[0]

	if len(pokemon) == 0 {
		return fmt.Errorf("No pokemon provided")
	}

	pokemondata, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemondata.Name)
	fmt.Println("Weight:", pokemondata.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemondata.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range pokemondata.Types {
		fmt.Printf("  -%v\n", poketype.Type.Name)
	}

	return nil
}
