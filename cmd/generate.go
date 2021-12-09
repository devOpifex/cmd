package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devOpifex/cmd/parser"
	"github.com/spf13/cobra"
)

var dir string
var config parser.Config

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate code for a command line application",
	Run: func(cmd *cobra.Command, args []string) {
		confFile := filepath.Join(dir, "cmd.json")

		_, err := os.Stat(confFile)

		if err != nil {
			fmt.Println("Cannot find cmd.json in directory")
			os.Exit(1)
		}

		config, err = parser.Read(confFile)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = config.Generate(dir)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.Flags().StringVarP(&dir, "dir", "d", ".", "path to root of package")
	rootCmd.AddCommand(generateCmd)
}
