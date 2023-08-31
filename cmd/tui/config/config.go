package tuiconfig

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	modelconfig "github.com/codosseum-org/terminal-client/internal/model/config"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list   list.Model
	config modelconfig.Config
	items  []list.Item
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	m.list.SetItems(m.items)
	return m.list.View()
}

func NewModel(config modelconfig.Config) Model {
	items := []list.Item{
		item{"Host", config.General.URL},
		item{"Default Name", config.General.DefaultName},
	}

	m := Model{
		list:   list.New(items, list.NewDefaultDelegate(), 0, 0),
		config: config,
		items:  items,
	}
	m.list.SetShowHelp(true)
	m.list.Title = "Settings"

	return m
}
