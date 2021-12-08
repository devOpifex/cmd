package parser

import (
	"errors"
	"os"
	"path/filepath"
)

var code string = `package main

import (
	"github.com/spf13/cobra"
)

func main() {
}`

func (conf Config) setup(dir string) error {
	programDir := filepath.Join(dir, conf.Program)

	_, err := os.Stat(programDir)

	if errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(programDir, 0755)
	}

	return nil
}

func (conf Config) Generate(dir string) error {
	err := conf.setup(dir)

	if err != nil {
		return err
	}

	mainFile := filepath.Join(dir, conf.Program, "main.go")

	f, err := os.Create(mainFile)

	if err != nil {
		return errors.New("could not create main.go file")
	}

	f.WriteString(code)

	return nil
}
