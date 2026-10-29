package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	mon := args[0]
	if p, exists := cfg.pokedex[mon]; exists {
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Height: %d\n", p.Height)
		fmt.Printf("Weight: %d\n", p.Weight)
		fmt.Printf("Stats:\n")
		for _, s := range p.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

	} else {
		fmt.Printf("You haven't caught %s\n", mon)
	}
	return nil
}
