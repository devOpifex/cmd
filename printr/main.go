package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
		Use:   "printr",
		Short: "print stuff",
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
	func main() {
		execute()
	}
	var printCmd = &cobra.Command{
		Use:   "print",
		Short: "Print Something",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	