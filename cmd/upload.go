package cmd

import (
	"fmt"
	"os"

	"github.com/codosseum-org/terminal-client/pkg/upload"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file with your solution.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		err := upload.UploadAndValidateFile(filePath)
        if err != nil {
			fmt.Printf("Error: %v", err)
            os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
