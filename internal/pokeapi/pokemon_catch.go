package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonCatch(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	data, ok := c.pokeCache.Get(url)
	if ok {
		PokemonCache := Pokemon{}
		err := json.Unmarshal(data, &PokemonCache)
		if err != nil {
			return Pokemon{}, err
		}

		return PokemonCache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad response status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.pokeCache.Add(url, jsonData)

	pokemonResp := Pokemon{}
	err = json.Unmarshal(jsonData, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
