package tui

import (
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	FileName      string
	Content       string
	CompiledRegex *regexp.Regexp
	Err           error
	TextInput     textinput.Model
	Viewport      viewport.Model
	WindowHeight  int
	WindowWidth   int
}

var defaultBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("30")).
	Padding(0, 1)

func InitialModel(filepath, content string) Model {
	ti := textinput.New()
	ti.Placeholder = "Type regex expression here..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	
	vp := viewport.New(0, 0)
	vp.SetContent(content)

	return Model{
		FileName:  filepath,
		Content:   content,
		Err:       nil,
		TextInput: ti,
		Viewport:  vp,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
