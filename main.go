package main

func main() {
	startupURL := "https://pokeapi.co/api/v2/pokemon/"
	cfg := &config{
		commands: getCommands(),
		next:     &startupURL,
		previous: nil,
	}
	startRepl(cfg)
}
