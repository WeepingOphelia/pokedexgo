package main

import "fmt"

func commandHelp(cfg *config) error {

	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
