package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Program     string   `json:"program"`
	Package     string   `json:"package"`
	Description string   `json:"description"`
	Commands    Commands `json:"commands"`
}

type Commands []Command

type Command struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Function    string    `json:"function"`
	Arguments   Arguments `json:"arguments"`
}

type Arguments []Argument

type Argument struct {
	Name        string `json:"name"`
	Short       string `json:"short"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

func Read(path string) (Config, error) {
	var conf Config
	reader, err := os.ReadFile(path)

	if err != nil {
		return conf, err
	}

	json.Unmarshal(reader, &conf)

	return conf, nil
}

func (conf *Config) Check() error {
	if conf.Package == "" {
		return errors.New("config missing package")
	}

	if conf.Program == "" {
		return errors.New("config missing program")
	}

	forbiddenPrograms := []string{"inst", "R", "src"}
	for _, p := range forbiddenPrograms {
		if conf.Program == p {
			return fmt.Errorf("cannot use %v as program name", p)
		}
	}

	if len(conf.Commands) == 0 {
		return errors.New("config has no commands")
	}

	return conf.Commands.check()
}

func (cmds Commands) check() error {
	for i, cmd := range cmds {
		if len(cmd.Arguments) == 0 {
			return fmt.Errorf("command %v: missing arguments", i)
		}

		if cmd.Description == "" {
			return fmt.Errorf("command %v: missing description", i)
		}

		if cmd.Function == "" {
			return fmt.Errorf("command %v: missing function", i)
		}

		if cmd.Name == "" {
			return fmt.Errorf("command %v: missing name", i)
		}
	}
	return nil
}

func parseType(input string) string {
	switch input {
	case "character":
		return "string"
	case "numeric":
		return "float64"
	case "integer":
		return "int"
	default:
		return input
	}
}

func parseDefault(s, t string) string {
	if t == "string" {
		return "\"" + s + "\""
	}

	return t
}
