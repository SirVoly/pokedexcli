package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetLocationsList(pageURL *string) (PokeAPILocationResponse, error) {
	locationResponse := PokeAPILocationResponse{}

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return locationResponse, err
		}

		return locationResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationResponse, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResponse, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse, err
	}

	err = json.Unmarshal(dat, &locationResponse)
	if err != nil {
		return locationResponse, err
	}

	c.cache.Add(url, dat)

	return locationResponse, nil
}

func (c *Client) GetPokemonsAtLocationList(area string) (LocationArea, error) {
	var location LocationArea
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &location)
		if err != nil {
			return location, err
		}

		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return location, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return location, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(dat, &location)
	if err != nil {
		return location, err
	}

	c.cache.Add(url, dat)

	return location, nil
}
