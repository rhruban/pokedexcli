package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name
	
	if val, ok := c.cache.Get(url); ok {
		poks := RespPokemon{}
		err := json.Unmarshal(val, &poks)
		if err != nil {
			return RespPokemon{}, err
		}
		return poks, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}
	
	poks := RespPokemon{}
	err = json.Unmarshal(dat, &poks)
	if err != nil {
		return RespPokemon{}, err
	}
	
	c.cache.Add(url, dat)
	return poks, nil
}


