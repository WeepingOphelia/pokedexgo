package main

import (
	"fmt"

	"github.com/WeepingOphelia/pokedexgo/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	mapArea, err := pokeapi.RetrieveMap(cfg.mapNext, cfg.mapCache)
	if err != nil {
		return err
	}

	cfg.mapNext = mapArea.Next
	cfg.mapPrev = mapArea.Previous

	for _, result := range mapArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}
