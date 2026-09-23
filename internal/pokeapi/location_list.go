package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.pokeCache.Get(url)
	if ok {
		locationsCache := RespShallowLocations{}
		err := json.Unmarshal(data, &locationsCache)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsCache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if resp.StatusCode > 299 {
		return RespShallowLocations{}, fmt.Errorf("bad response status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.pokeCache.Add(url, jsonData)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(jsonData, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
