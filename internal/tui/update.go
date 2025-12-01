package tui

import (
	"fmt"
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.WindowHeight = msg.Height
		m.WindowWidth = msg.Width
		// for the TextInput (1 line) + some margin (2 lines in this case) + spearator (1 line)
		headerHeight := 5
		m.Viewport.Height = m.WindowHeight - headerHeight
		m.Viewport.Width = m.WindowWidth
		m.TextInput.Width = m.WindowWidth - 5
		m.Viewport, cmd = m.Viewport.Update(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.Viewport.ScrollUp(1)
		case "down", "j":
			m.Viewport.ScrollDown(1)
		case "esc":
			if m.TextInput.Focused() {
				m.TextInput.Blur()
			}

		default:
			if !m.TextInput.Focused() && msg.Type == tea.KeyRunes {
				cmd = m.TextInput.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case errMsg:
		m.Err = msg
		return m, nil
	}

	newTextInput, tiCmd := m.TextInput.Update(msg)
	cmds = append(cmds, tiCmd)

	if newTextInput.Value() != m.TextInput.Value() {
		m.TextInput = newTextInput
		pattern := m.TextInput.Value()

		if pattern == "" {
			m.CompiledRegex = nil
			m.Err = nil
		} else {
			re, err := regexp.Compile(pattern)
			if err != nil {
				m.Err = fmt.Errorf("invalid regex expression %w", err)
				m.CompiledRegex = nil
			} else {
				m.CompiledRegex = re
				m.Err = nil
			}
		}
	} else {
		m.TextInput = newTextInput
	}

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
