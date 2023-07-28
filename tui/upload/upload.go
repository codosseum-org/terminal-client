package upload

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codosseum-org/terminal-client/pkg/api"
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
	spinner  spinner.Model
	ready    bool
	uploaded bool
}

func NewModel(data, name, language string) Model {
	ta := textarea.New()
	ta.Placeholder = "It's empty here..."
	ta.MaxHeight = 0
	ta.MaxWidth = 0
	ta.CharLimit = 0

	s := spinner.New()
	s.Spinner = spinner.MiniDot

	return Model{
		data:     data,
		fileName: name,
		fileLang: language,
		keys:     keys,
		help:     help.New(),
		spinner:  s,
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
				err := api.UploadCode(m.textarea.Value())
				if err != nil {
					log.Fatal("Okay yeah this can't even happen right now!")
				}
				m.uploaded = true
                return m, m.spinner.Tick
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
	var s string
	s += m.textarea.View()

	fileInfo := "Reviewing " + m.fileName + " | Language: " + m.fileLang
    s += "\n" + fileInfo
    if m.uploaded {
        s += " | " + m.spinner.View() + " Uploading (to be done, quit for now!)"
    }
    s += "\n" + m.helpView()
    return s
}
