package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/abdul-rehman-d/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

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
