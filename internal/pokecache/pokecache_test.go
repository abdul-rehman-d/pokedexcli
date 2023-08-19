package pokecache

import (
	"testing"
)

func TestPokeCache(t *testing.T) {
	cache := NewCache()

	if cache.cache == nil {
		t.Error("cache should be a map, not nil")
	}

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
	}

	for _, c := range cases {

		cache.Add(c.inputKey, c.inputVal)

		actual, exists := cache.Get(c.inputKey)

		if !exists {
			t.Errorf("%v should exists on cache", c.inputKey)
			continue
		}

		if string(actual) != string(c.inputVal) {
			t.Errorf("%v is not equal to %v", string(actual), string(c.inputVal))
			continue
		}
	}
}
