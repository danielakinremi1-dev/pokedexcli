package main

import ("fmt"
		"net/http")

type location struct { 
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}



func commandMap(cfg *config, page string) error {
	if page == "map"{
		return 
	} elif page == "mapb" {
		return
	}
	
	

	for _, cmd := range cfg.commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
