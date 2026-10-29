package main

import (
	"fmt"
	"math/rand"

	"github.com/WeepingOphelia/pokedexgo/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	mon := args[0]
	monStats, err := pokeapi.QueryMon(mon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", mon)
	n := rand.Intn(400)
	if n > monStats.BaseExperience {
		fmt.Printf("%s was caught!\n", mon)
		cfg.pokedex[mon] = monStats
	} else {
		fmt.Printf("%s escaped!\n", mon)
	}
	return nil
}
