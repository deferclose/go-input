package main

import (
	"log"

	"github.com/deferclose/go-input"
)

func main() {

	ui := &input.UI{}

	query := "What is your name?"
	name, err := ui.Ask(query, &input.Options{
		// Read the default val from env var
		Default:  "alan",
		Required: true,
		Loop:     true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer is %s\n", name)
}
