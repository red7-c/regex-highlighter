package tui

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e1e2e3ff")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#005ae1ff"))

	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#696969ff")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#6492d7ff"))

	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0000")).
			Padding(0, 1)

	counterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(0, 1)

	HighlightStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#ffae5cff")).
			Foreground(lipgloss.Color("0"))

	SelectedMatchStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#0230f8ff")).
				Foreground(lipgloss.Color("#f3f1efff"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(0, 1)

	sideBarStyles = lipgloss.NewStyle().
			Width(30).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 1).
			MarginLeft(1)
)
