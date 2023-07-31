package cmd

import (
	"fmt"
	"log"

	pkgconfig "github.com/codosseum-org/terminal-client/pkg/tui/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Opens a configuration TUI that lets you modify the behaviour of the client.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world :)")
		err := pkgconfig.StartTUI()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// generate config as well
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
