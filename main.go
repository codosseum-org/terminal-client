package main

// Should also import premade components from https://github.com/charmbracelet/bubbles, as well as https://github.com/charmbracelet/lipgloss for styling.
import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

// Intiialize model
func initialModel() model {
	return model{}
}

// Initialize BubbleTea
func (m model) Init() tea.Cmd {
	return nil
}

// When an event happens
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        }
    }
	return m, nil
}

// UI Setup
func (m model) View() string {
	ui := "Hi hello, press 'q' to exit out"

	return ui
}
