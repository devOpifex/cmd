package parser

import (
	"errors"
	"os"
	"path/filepath"
)

type app struct {
	code string
	conf Config
}

var main string = `package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
`

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

	defer f.Close()

	var cli = app{
		code: main,
		conf: conf,
	}
	cli.root()
	cli.mainFn()

	f.WriteString(cli.code)

	return nil
}

func (main *app) root() {
	var root = `var rootCmd = &cobra.Command{
		Use:   "` + main.conf.Program + `",
		Short: "` + main.conf.Description + `",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	func execute() {
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	`
	main.code = main.code + root
}

func (main *app) mainFn() {
	main.code = main.code + `func main() {
		execute()
	}
	`
}
