package pokeapi

import(
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string) LocationAreas {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	
	locs := LocationAreas{}
	err = json.Unmarshal(body, &locs)
	if err != nil {
		fmt.Println(err)
	}
	return locs
}
