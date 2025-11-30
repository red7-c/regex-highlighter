package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f7f7f7ff")).Render
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

var (
	TitleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	InfoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return TitleStyle.BorderStyle(b)
	}()

	validInputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#444444ff")).
			Padding(0, 1)

	errorInputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF0000")).
			Padding(0, 1)
)

func (m Model) HeaderView() string {
	title := TitleStyle.Render(m.fileName)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) FooterView() string {
	info := InfoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s",
		helpStyle(m.HeaderView()),
		textStyle(m.viewport.View()),
		validInputStyle.Render(m.textInput.View()),
	)
}
