package pokeapi

import (
	"encoding/json"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func getPaginatedList[T any](c *Client, url string) (PokeAPIPageResponse[T], error) {
	var pokeAPIPageResponse PokeAPIPageResponse[T]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeAPIPageResponse, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokeAPIPageResponse, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pokeAPIPageResponse); err != nil {
		return pokeAPIPageResponse, err
	}

	return pokeAPIPageResponse, nil
}

func (c *Client) GetLocationsList(pageURL *string) (PokeAPIPageResponse[LocationArea], error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	return getPaginatedList[LocationArea](c, url)
}
