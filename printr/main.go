package main

import (
	"bufio"
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
	var printStr string
func main() {
		execute()
	}
	var printCmd = &cobra.Command{
		Use: "print",
		Short: "Print Something",
		Run: func(cmd *cobra.Command, args []string) {
			rArgs := "-e 'printer::print_sth("+printStr+")'"
			path := getPath()
			rCommand := exec.Command(path, rArgs)
			stdout, err := rCommand.StdoutPipe()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			rCommand.Start()
			scanner := bufio.NewScanner(stdout)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}
			rCommand.Wait()	
		},
	}
	func init(){
printCmd.Flags().StringVarP(&printStr,"str","s","hello world","string to print")
}
