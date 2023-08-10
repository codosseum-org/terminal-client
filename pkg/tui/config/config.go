package pkgconfig

import (
	tea "github.com/charmbracelet/bubbletea"
	tuiconfig "github.com/codosseum-org/terminal-client/cmd/tui/config"
	"github.com/codosseum-org/terminal-client/internal/config"
)

func StartTUI() error {
	config, err := config.GetConfig()
	if err != nil {
		return err
	}
	p := tea.NewProgram(tuiconfig.NewModel(config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
