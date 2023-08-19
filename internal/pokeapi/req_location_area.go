package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	var endpoint string
	if pageUrl == nil {
		endpoint = baseUrl + "/location-area"
	} else {
		endpoint = *pageUrl
	}

	data, err := c.request(endpoint)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	var location LocationAreasResponse

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return location, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	var endpoint string
	endpoint = baseUrl + "/location-area/" + locationAreaName

	data, err := c.request(endpoint)

	if err != nil {
		return LocationArea{}, err
	}

	var location LocationArea

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}
