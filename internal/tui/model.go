package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	fileName  string
	Content   string
	ready     bool
	Err       error
	textInput textinput.Model
	viewport  viewport.Model
}

func InitialModel(filepath, content string) Model {

	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#acaca8ff"))

	ti := textinput.New()
	ti.Placeholder = "Type regex expression here..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 60
	ti.Cursor.Style = cursorStyle

	return Model{
		fileName:  filepath,
		Content:   content,
		Err:       nil,
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
