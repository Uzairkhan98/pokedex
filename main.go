package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var command map[string]cliCommand
var exit = false

func init() {
	command = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage: \n\n%v: %v\n%v: %v\n\n", command["help"].name, command["help"].description, command["exit"].name, command["exit"].description)
	return nil
}

func commandExit() error {
	exit = true
	return nil
}


func main() {
	reader := bufio.NewReader(os.Stdin)

    for !exit{
        fmt.Print("Pokedex > ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        input = strings.TrimSpace(input)
		if command, ok := command[input]; ok {
			err := command.callback()
            if err!= nil {
                fmt.Println("Error executing command:", err)
                continue
            }
		}
    }
}