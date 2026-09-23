package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon)
	}

	return nil
}
