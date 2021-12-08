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

var rootCmd = &cobra.Command{
	Use:   "rcmd",
	Short: "Make command line program",
	Long:  `Make Command line program`,
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

		err = config.Generate()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "path to root of package")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
