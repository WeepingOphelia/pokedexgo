package main

import (
	"fmt"

	"github.com/WeepingOphelia/pokedexgo/internal/pokeapi"
)

func commandMapBack(cfg *config, args []string) error {
	if cfg.mapPrev == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	mapArea, err := pokeapi.RetrieveMap(*cfg.mapPrev, cfg.mapCache)
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
