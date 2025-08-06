package pokeapi

import (
	"io"
	"net/http"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetLocationsList(url string) ([]byte, error) {
	var dat []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dat, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return dat, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return dat, err
	}

	return dat, nil
}
