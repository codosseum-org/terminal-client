package upload

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Upload     key.Binding
	ChangeMode key.Binding
	Quit       key.Binding
    Tab        key.Binding
}

var (
	keys = keyMap{
		Upload: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "upload & submit file"),
		),
		ChangeMode: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "change modes"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
        Tab: key.NewBinding(
            key.WithKeys("tab"),
            key.WithHelp("tab", "tab"),
        ),
	}
	normalKeymap = []key.Binding{
		keys.Upload, keys.ChangeMode, keys.Quit,
	}
	focusedKeymap = []key.Binding{
		keys.ChangeMode, keys.Quit,
	}
)

type Model struct {
	data     string
	fileName string
	fileLang string
	keys     keyMap
	help     help.Model
	textarea textarea.Model
	ready    bool
}

func NewModel(data, name, language string) Model {
	ta := textarea.New()
	ta.Placeholder = "It's empty here..."
	ta.MaxHeight = 0
	ta.MaxWidth = 0
	ta.CharLimit = 0

	return Model{
		data:     data,
		fileName: name,
		fileLang: language,
		keys:     keys,
		help:     help.New(),
		textarea: ta,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

		m.textarea.SetWidth(msg.Width)
		m.textarea.SetHeight(msg.Height - 2)
		if !m.ready {
			m.textarea.SetValue(m.data)
		}
	case tea.KeyMsg:
		if !m.textarea.Focused() {
			switch {
			case key.Matches(msg, m.keys.Upload):
				fmt.Print("Upload logic!")
			case key.Matches(msg, m.keys.ChangeMode):
				m.textarea.Focus()
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			}
		} else {
			switch {
			case key.Matches(msg, m.keys.ChangeMode):
				m.textarea.Blur()
            case key.Matches(msg, m.keys.Tab):
                m.textarea.InsertString("   ")
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			}
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) helpView() string {
	if m.textarea.Focused() {
		return m.help.ShortHelpView(focusedKeymap)
	} else {
		return m.help.ShortHelpView(normalKeymap)
	}
}

func (m Model) View() string {
    fileInfo := "Reviewing " + m.fileName + " | Language: " + m.fileLang
	return m.textarea.View() + "\n" + fileInfo + "\n" + m.helpView()
}
