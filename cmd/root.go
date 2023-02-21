package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "[command]",
	Short: "File Operations",
	Long:  `File Operations, including size, create, delete ...`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sizeCmd, completionCmd) // adding add command to root
}
