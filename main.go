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

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
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
	fmt.Printf("\nWelcome to the Pokedex!\nUsage: \n\n")
	for _, command := range commands {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}


func main() {
	reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        input = strings.TrimSpace(input)
		if command, ok := commands[input]; ok {
			err := command.callback()
            if err!= nil {
                fmt.Println("Error executing command:", err)
                continue
            }
		}
    }
}