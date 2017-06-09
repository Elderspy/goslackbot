package main

import (
	"errors"
	"fmt"
	"github.com/Elderspy/goslackbot/helpers"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("dicks")
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
			if parts[0] == "butts" {
				m.Text = "I like big butts and I can not lie."
				helpers.SendMessage(ws, m)
			}
		}

	}
}
