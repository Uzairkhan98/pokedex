package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Config
	Count    int    `json:"count"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

func pokeapi(url string, con *Config) (Location, error)  {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return  Location{}, err
	}

    // Unmarshal the JSON data into the Person struct
    var locations Location
    err = json.Unmarshal([]byte(body), &locations)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return Location{}, err
    }
	if locations.Next != nil {
		*con.Next = *locations.Next
	} else {
		*con.Next = ""
	}
	if locations.Previous != nil {
		*con.Previous = *locations.Previous
	} else {
		*con.Previous = ""
	}
	return locations, err
}

