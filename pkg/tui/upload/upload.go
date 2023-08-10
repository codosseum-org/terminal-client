package pkgupload

import (
	tea "github.com/charmbracelet/bubbletea"
	tuiupload "github.com/codosseum-org/terminal-client/cmd/tui/upload"
	"github.com/codosseum-org/terminal-client/internal/upload"
)

func StartTUI(path string) error {
    data, err := upload.GetFileContent(path)
	if err != nil {
		return err
	}

    name, language := upload.GetFileInformation(path)
    p := tea.NewProgram(tuiupload.NewModel(string(data), name, language), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        return err
    }

    return nil
}

func UploadCode(code string) error {
    

    return nil
}
