package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "codosseum",
    Short: "Official Codosseum terminal client TUI & CLI",
    Long: "Official Codosseum terminal client TUI & CLI",

    Run: func(cmd *cobra.Command, args []string) {

    },
}


func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
