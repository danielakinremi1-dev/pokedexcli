package main

import (
	"time"

	"github.com/danielakinremi1-dev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: &pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
