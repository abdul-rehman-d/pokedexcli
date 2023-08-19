package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) request(endpoint string) ([]byte, error) {
	var data []byte

	cachedRes, exists := c.cache.Get(endpoint)

	if exists {
		data = cachedRes
		fmt.Println("cache hit")
	} else {
		fmt.Println("cache miss")
		req, err := http.NewRequest("GET", endpoint, nil)

		if err != nil {
			return []byte{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return []byte{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 399 {
			return []byte{}, fmt.Errorf("bad request: %v", res.StatusCode)
		}

		httpRes, err := io.ReadAll(res.Body)

		if err != nil {
			return []byte{}, err
		}

		data = httpRes
		c.cache.Add(endpoint, data)
	}

	return data, nil
}

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
	endpoint = baseUrl + "/location-area"

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
