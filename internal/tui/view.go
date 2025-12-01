package tui

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205"))

	blurredStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240"))

	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	HighlightStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#ecdc20")).
			Foreground(lipgloss.Color("0"))
)

func (m Model) ApplyRegexHighlighting(content, pattern string) string {
	if pattern == "" {
		return content
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "Error: invalid regex espression" + err.Error()
	}

	normalizedContent := strings.ReplaceAll(m.Content, "\r\n", "\n")
	lines := strings.Split(normalizedContent, "\n")

	var styledLines []string

	for _, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		styledLine := m.ReconstructLine(line, matches)
		styledLines = append(styledLines, styledLine)
	}

	return strings.Join(styledLines, "\n")
}

func (m Model) ReconstructLine(rawLine string, matches [][]int) string {

	style := HighlightStyle
	var b strings.Builder
	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			b.WriteString(rawLine[lastIndex:start])
		}

		matchedText := rawLine[start:end]
		b.WriteString(style.Render(matchedText))

		lastIndex = end
	}

	if lastIndex < len(rawLine) {
		b.WriteString(rawLine[lastIndex:])
	}

	return b.String()
}

func (m Model) View() string {

	var inputView string
	if m.TextInput.Focused() {
		inputView = focusedStyle.Render(m.TextInput.View())
	} else {
		inputView = blurredStyle.Render(m.TextInput.View())
	}
	pattern := m.TextInput.Value()
	highlightedContent := m.ApplyRegexHighlighting(m.Content, pattern)
	m.Viewport.SetContent(highlightedContent)

	ui := strings.Join([]string{
		inputView,
		separatorStyle.Render(strings.Repeat("—", m.WindowWidth)),
		m.Viewport.View(),
	}, "\n")

	return ui
}
