package main

import (
	"fmt"

	"github.com/WeepingOphelia/pokedexgo/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location to explore")
	}

	fmt.Printf("Exploring %s...\n", args[0])

	availableMons, err := pokeapi.ExploreArea(args[0], cfg.mapCache)
	if err != nil {
		return fmt.Errorf("error exploring area: %w", err)
	}

	if len(availableMons) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return nil
	}

	fmt.Printf("Found %d Pokémon:\n", len(availableMons))

	for _, mon := range availableMons {
		fmt.Println("# " + mon)
	}

	return nil
}
