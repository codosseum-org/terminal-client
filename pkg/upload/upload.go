package upload

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codosseum-org/terminal-client/tui/upload"
)

// Must add real uploading logic when the backend is complete!
func UploadAndValidateFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

    p := tea.NewProgram(upload.NewModel(string(data)), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        return err
    }

    return nil
}
