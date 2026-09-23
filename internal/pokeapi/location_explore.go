package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(location string) (RespDetailedLocation, error) {
	url := baseURL + "/location-area/" + location

	data, ok := c.pokeCache.Get(url)
	if ok {
		locationCache := RespDetailedLocation{}
		err := json.Unmarshal(data, &locationCache)
		if err != nil {
			return RespDetailedLocation{}, err
		}

		return locationCache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	if resp.StatusCode > 299 {
		return RespDetailedLocation{}, fmt.Errorf("bad response status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocation{}, err
	}
	c.pokeCache.Add(url, jsonData)

	locationResp := RespDetailedLocation{}
	err = json.Unmarshal(jsonData, &locationResp)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	return locationResp, nil
}
