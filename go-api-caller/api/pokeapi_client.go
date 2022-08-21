package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mtslzr/pokeapi-go/structs"
)

const apiurl = "https://pokeapi.co/api/v2/pokemon/"

func buildHttpClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

func buildGetPokemonRequest(name string) (*http.Request, error) {
	url := apiurl + name
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return req, err
}

func get(req *http.Request) ([]byte, error) {
	client := buildHttpClient()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error Response. Status:%s, Body:%s", resp.Status, body)
	}

	return body, err
}

func GetPokemon(name string) (*structs.Pokemon, error) {
	req, err := buildGetPokemonRequest(name)
	if err != nil {
		return nil, err
	}

	body, err := get(req)
	if err != nil {
		return nil, err
	}

	poke := &structs.Pokemon{}
	err = json.Unmarshal(body, poke)
	if err != nil {
		return nil, err
	}
	return poke, nil
}
