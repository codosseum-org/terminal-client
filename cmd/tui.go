package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command {
    Use: "tui",
    Short: "Do TUI stuff!",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("You're in the TUI now ok?")
    },
}

func init() {
    rootCmd.AddCommand(tuiCmd)
}
