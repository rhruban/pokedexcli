package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	if val, ok := c.cache.Get(url); ok {
		locs := RespLocationAreas{}
		err := json.Unmarshal(val, &locs)
		if err != nil {
			return RespLocationAreas{}, err
		}
		return locs, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}
	
	locs := RespLocationAreas{}
	err = json.Unmarshal(dat, &locs)
	if err != nil {
		return RespLocationAreas{}, err
	}
	
	c.cache.Add(url, dat)
	return locs, nil
}
