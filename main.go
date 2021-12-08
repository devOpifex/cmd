package main

import (
	"log"

	"github.com/devOpifex/cmd/cmd"
	"github.com/devOpifex/cmd/parser"
)

func main() {
	parsed, err := parser.Read("cmd.json")

	if err != nil {
		log.Fatal(err)
	}

	err = parsed.Check()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
