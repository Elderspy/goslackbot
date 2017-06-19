package main

import (
	"errors"
	"fmt"
	"./helpers"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		panic(errors.New("You need to pass a token"))
	}

	token := os.Args[1]

	ws, id := helpers.SlackConnect(token)
	fmt.Println("This is our ID: " + id)
	for {
		m, err := helpers.GetMessage(ws)
		if err != nil {
			panic(err)
		}
		if m.Type == "message" {
			parts := strings.Fields(m.Text)
			fmt.Println(parts)
			if parts[0] == "test" {
				m.Text = "I don't do anything useful yet"
				helpers.SendMessage(ws, m)
			}
		}

	}
}
