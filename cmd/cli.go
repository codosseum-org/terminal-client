package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command {
    Use: "cli",
    Short: "Do CLI stuff!",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hi welcome to the CLI :)")
    },
}

func init() {
    rootCmd.AddCommand(cliCmd)
}
