package main

import "os"

func commandExit(con *Config) error {
	os.Exit(0)
	return nil
}
