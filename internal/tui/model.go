package tui

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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
	Matches       [][]int
	CurrentMatch  int
	ShowHelp      bool
}

func InitialModel(filepath, content string) Model {

	content = strings.ReplaceAll(content, "\r\n", "\n")

	ti := textinput.New()
	ti.Placeholder = "Type expression here..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 20

	vp := viewport.New(0, 0)
	vp.SetContent(content)

	return Model{
		FileName:  filepath,
		Content:   content,
		Err:       nil,
		TextInput: ti,
		Viewport:  vp,
		ShowHelp:  false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
