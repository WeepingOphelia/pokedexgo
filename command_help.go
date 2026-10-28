package main

import "fmt"

func commandHelp(cfg *config, args []string) error {

	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
