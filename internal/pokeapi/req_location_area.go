package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error) {
	endpoint := baseUrl + "/location"

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

	data, err := io.ReadAll(res.Body)

	var location LocationAreasResponse

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return location, nil
}
