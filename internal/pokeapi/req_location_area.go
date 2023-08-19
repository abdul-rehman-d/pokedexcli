package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	var endpoint string
	if pageUrl == nil {
		endpoint = baseUrl + "/location-area"
	} else {
		endpoint = *pageUrl
	}

	var data []byte

	cachedRes, exists := c.cache.Get(endpoint)

	if exists {
		data = cachedRes
		fmt.Println("cache hit")
	} else {
		fmt.Println("cache miss")
		req, err := http.NewRequest("GET", endpoint, nil)

		if err != nil {
			return LocationAreasResponse{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationAreasResponse{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 399 {
			return LocationAreasResponse{}, fmt.Errorf("bad request: %v", res.StatusCode)
		}

		httpRes, err := io.ReadAll(res.Body)

		if err != nil {
			return LocationAreasResponse{}, err
		}

		data = httpRes
		c.cache.Add(endpoint, data)
	}

	var location LocationAreasResponse

	err := json.Unmarshal(data, &location)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return location, nil
}
