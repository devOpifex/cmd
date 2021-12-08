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
	"os/exec"

	"github.com/spf13/cobra"
)

func getPath() string {
	var p string

	p, err := exec.LookPath("R")

	if err != nil {
		fmt.Println("Could not locate R installation")
	}

	return p
}
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
	cli.flags()

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

func (app *app) flags() {
	app.code = app.code + "func init(){\n"
	for _, cmd := range app.conf.Commands {
		for _, arg := range cmd.Arguments {
			def := parseDefault(arg.Default, arg.Type)
			app.code = app.code + cmd.Name + "Cmd.Flags()." + strings.Title(parseType(arg.Type)) + "VarP(&" + cmd.Name + strings.Title(arg.Name) + ",\"" + arg.Name + "\",\"" + arg.Short + "\"," + def + ",\"" + arg.Description + "\")\n"
		}
	}
	app.code = app.code + "}\n"
}

func (app *app) variables() {
	for _, cmd := range app.conf.Commands {
		for _, arg := range cmd.Arguments {
			app.code = app.code + "var " + cmd.Name + strings.Title(arg.Name) + " " + parseType(arg.Type) + "\n"
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
			path := getPath()
			rCommand := exec.Command(path, rArgs)
			stdout, _ := rCommand.StdoutPipe()

			rCommand.Start()
			if verbose {
				scanner := bufio.NewScanner(stdout)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					line := scanner.Text()
					fmt.Println(line)
				}
			}
			rCommand.Wait()	
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
