package main

import (
	"errors"
	"fmt"
)

func commandMapB(con *Config) error {
	if(len(*con.Previous) == 0) {
		return errors.New("cannot go back from the first page")
	}
	url := *con.Previous
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