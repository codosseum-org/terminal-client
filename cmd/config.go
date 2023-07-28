package cmd

import (
	"fmt"
	"log"

	"github.com/codosseum-org/terminal-client/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Opens a configuration TUI that lets you modify the behaviour of the client.",
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello world :)")
        err := config.GenerateConfig()
        if err != nil {
            log.Fatalf("Error: %v", err)
        }
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
