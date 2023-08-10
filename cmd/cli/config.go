package cmd

import (
	"log"

	"github.com/codosseum-org/terminal-client/internal/config"
	pkgconfig "github.com/codosseum-org/terminal-client/pkg/tui/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Opens a configuration TUI that lets you modify the behaviour of the client.",
	Run: func(cmd *cobra.Command, args []string) {
		if !config.DoesConfigExist() {
			err := config.GenerateConfig()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}
		err := pkgconfig.StartTUI()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
