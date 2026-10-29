package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func QueryMon(mon string) (Pokemon, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + mon)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var p Pokemon
	err = json.Unmarshal(dat, &p)
	if err != nil {
		return Pokemon{}, err
	}

	return p, nil
}
