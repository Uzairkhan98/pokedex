package main

import "fmt"

func commandMap(con *Config) error {
	url := "https://pokeapi.co/api/v2/location/"
	if len(*con.Next) > 0 {
		url = *con.Next
	}
	locations, err := pokeapi(url, con)
	if err != nil {
		return err
	}
	// Display the locations
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}