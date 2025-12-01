package tui

import (
	"fmt"
	"regexp"
	"strings"

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
		m.updateSizes()
		m.Viewport, cmd = m.Viewport.Update(msg)
		return m, nil

	case tea.KeyMsg:
		if m.ShowHelp {
			if msg.String() == " " || msg.String() == "esc" || msg.String() == "q" {
				m.ShowHelp = false
				m.updateSizes()
			}
			return m, nil
		}
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

		case " ":
			if !m.TextInput.Focused() {
				m.ShowHelp = true
				m.updateSizes()
				return m, nil
			}

		case "n":
			if !m.TextInput.Focused() && len(m.Matches) > 0 {
				m.CurrentMatch = (m.CurrentMatch + 1) % len(m.Matches)
				m.scrollToMatch()
			}
		case "p":
			if !m.TextInput.Focused() && len(m.Matches) > 0 {
				m.CurrentMatch = (m.CurrentMatch - 1 + len(m.Matches)) % len(m.Matches)
				m.scrollToMatch()
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
			m.Matches = nil
			m.CurrentMatch = 0
		} else {
			re, err := regexp.Compile(pattern)
			if err != nil {
				m.Err = fmt.Errorf("invalid regex expression: %w", err)
				m.CompiledRegex = nil
				m.Matches = nil
				m.CurrentMatch = 0
			} else {
				m.CompiledRegex = re
				m.Err = nil
				m.Matches = re.FindAllStringIndex(m.Content, -1)
				m.CurrentMatch = 0
			}
		}
		m.updateSizes()
	} else {
		m.TextInput = newTextInput
	}

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) scrollToMatch() {
	if len(m.Matches) == 0 {
		return
	}
	match := m.Matches[m.CurrentMatch]
	lineNum := strings.Count(m.Content[:match[0]], "\n")
	m.Viewport.SetYOffset(lineNum)
}

func (m *Model) updateSizes() {
	headerHeight := 5
	if m.Err != nil {
		headerHeight += 2
	}

	helpWidth := 0
	if m.ShowHelp {
		helpWidth = 35
	}

	availableWidth := m.WindowWidth - helpWidth
	if availableWidth < 0 {
		availableWidth = 0
	}

	m.Viewport.Height = m.WindowHeight - headerHeight
	m.Viewport.Width = availableWidth
	m.TextInput.Width = availableWidth - 5
}

func (m *Model) updateViewportHeight() {
	m.updateSizes()
}
