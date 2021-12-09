package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rcmd",
	Short: "Make command line program",
	Long:  `Make Command line program`,
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
