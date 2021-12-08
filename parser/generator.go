package parser

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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
	cli.variables()
	cli.execute()
	cli.cmds()

	f.WriteString(cli.code)

	return nil
}

func (app *app) root() {
	var root = `var rootCmd = &cobra.Command{
		Use:   "` + app.conf.Program + `",
		Short: "` + app.conf.Description + `",
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
	app.code = app.code + root
}

func (app *app) variables() {
	for _, cmd := range app.conf.Commands {
		for _, variable := range cmd.Arguments {
			app.code = app.code + "var " + cmd.Name + strings.Title(variable.Name) + " " + parseType(variable.Type) + "\n"
		}
	}
}

func (app *app) cmds() {
	for i := range app.conf.Commands {
		app.cmd(i)
	}
}

func (app *app) cmd(index int) {
	app.code = app.code + `var ` + app.conf.Commands[index].Name + `Cmd = &cobra.Command{
		Use:   "` + app.conf.Commands[index].Name + `",
		Short: "` + app.conf.Commands[index].Description + `",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	`
}

func (app *app) execute() {
	app.code = app.code + `func main() {
		execute()
	}
	`
}
