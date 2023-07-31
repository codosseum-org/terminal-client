package cmd

import (
	"log"

	pkgupload "github.com/codosseum-org/terminal-client/pkg/tui/upload"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file with your solution.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		err := pkgupload.StartTUI(filePath)
        if err != nil {
            log.Fatalf("Error: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
