package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationPokemon(area string) (RespLocationDetails, error) {
	url := baseURL + "/location-area/" + area
	
	if val, ok := c.cache.Get(url); ok {
		locs := RespLocationDetails{}
		err := json.Unmarshal(val, &locs)
		if err != nil {
			return RespLocationDetails{}, err
		}
		return locs, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetails{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationDetails{}, err
	}
	
	locs := RespLocationDetails{}
	err = json.Unmarshal(dat, &locs)
	if err != nil {
		return RespLocationDetails{}, err
	}
	
	c.cache.Add(url, dat)
	return locs, nil
}


