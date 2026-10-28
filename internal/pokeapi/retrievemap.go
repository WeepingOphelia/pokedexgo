package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/WeepingOphelia/pokedexgo/internal/pokecache"
)

func RetrieveMap(url string, cache *pokecache.Cache) (MapArea, error) {
	if cache != nil {
		if cached, ok := cache.Get(url); ok {
			mapArea := MapArea{}
			err := json.Unmarshal(cached, &mapArea)
			if err != nil {
				return MapArea{}, err
			}
			return mapArea, nil
		}
	}
	res, err := http.Get(url)
	if err != nil {
		return MapArea{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return MapArea{}, err
	}

	mapArea := MapArea{}
	err = json.Unmarshal(dat, &mapArea)
	if err != nil {
		return MapArea{}, err
	}

	cache.Add(url, dat)
	return mapArea, nil
}
