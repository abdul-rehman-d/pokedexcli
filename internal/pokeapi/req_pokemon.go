package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	var endpoint string
	endpoint = baseUrl + "/pokemon/" + pokemonName

	data, err := c.request(endpoint)

	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
