package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "size operator",
	Long:  `size operator, get the given file size`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("file path: " + args[0])

		fInfo, err := os.Stat(args[0])

		if err != nil {
			fmt.Println(err)
		} else {
			fsize := fInfo.Size()
			fmt.Printf("The file size is %d bytes\n", fsize)
		}

	},
}
