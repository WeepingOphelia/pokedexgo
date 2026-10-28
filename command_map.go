package main


func commandMap(cfg *config) error {
	for i := cfg.loc; i < 20 + config.loc; i++ {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		
		fmt.Printlin(res.name)
	}