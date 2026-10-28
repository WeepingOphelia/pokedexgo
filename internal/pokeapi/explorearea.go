package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/WeepingOphelia/pokedexgo/internal/pokecache"
)

func ExploreArea(location string, cache *pokecache.Cache) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location

	if cache != nil {
		if cached, ok := cache.Get(url); ok {
			locData := LocationData{}
			err := json.Unmarshal(cached, &locData)
			if err != nil {
				return []string{}, err
			}
			return getMons(locData), nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	locData := LocationData{}
	err = json.Unmarshal(dat, &locData)
	if err != nil {
		return []string{}, err
	}

	cache.Add(url, dat)

	return getMons(locData), nil
}

func getMons(locData LocationData) []string {
	availableMons := []string{}

	for _, mon := range locData.PokemonEncounters {
		availableMons = append(availableMons, mon.Pokemon.Name)
	}

	return availableMons
}
