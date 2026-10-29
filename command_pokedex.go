package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	for p, _ := range cfg.pokedex {
		fmt.Printf("- %s\n", p)
	}
	return nil
}
