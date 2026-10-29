package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/WeepingOphelia/pokedexgo/internal/pokeapi"
	"github.com/WeepingOphelia/pokedexgo/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commands map[string]cliCommand
	mapNext  string
	mapPrev  *string
	mapCache *pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
}

func cleanInput(input string) []string {

	return strings.Fields(strings.ToLower(input))
}

func getConfig() *config {
	return &config{
		commands: map[string]cliCommand{
			"catch": {
				name:        "catch",
				description: "Attempt to catch a Pokemon",
				callback:    commandCatch,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"explore": {
				name:        "explore",
				description: "Explore a location",
				callback:    commandExplore,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a Pokemon in your Pokedex",
				callback:    commandInspect,
			},
			"map": {
				name:        "map",
				description: "Lists locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Lists previous locations",
				callback:    commandMapBack,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Lists all Pokemon in your Pokedex",
				callback:    commandPokedex,
			},
		},
		mapNext:  "https://pokeapi.co/api/v2/location-area/",
		mapPrev:  nil,
		mapCache: pokecache.NewCache(5 * time.Second),
		pokedex:  make(map[string]pokeapi.Pokemon),
	}
}

func pokeScan(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		cmd := words[0]

		if c, exists := cfg.commands[cmd]; exists {
			err := c.callback(cfg, words[1:])
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", cmd)
		}

	}
}
