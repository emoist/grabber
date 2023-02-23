package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	size_kb = 1 << 10 // 1024 bytes
	size_mb = 1 << 20 // 1048576 bytes
	size_gb = 1 << 30 // 1073741824 bytes
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
			return
		}

		size_b := fInfo.Size() // size in bytes

		if size_b < size_kb {
			fmt.Printf("The file size is %d bytes\n", size_b)
		} else if size_b < size_mb {
			fmt.Printf("The file size is %dK\n", size_b/size_kb)
		} else if size_b < size_gb {
			fmt.Printf("The file size is %dM\n", size_b/size_mb)
		} else {
			fmt.Printf("The file size is %dG\n", size_b/size_gb)
		}

	},
}
