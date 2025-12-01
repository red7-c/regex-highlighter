package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const helpContent = `
Regex Cheat Sheet

.       Any char except newline
*       Zero or more
+       One or more
?       Zero or one
^       Start of line
$       End of line
\d      Digit (0-9)
\w      Word char
\s      Whitespace
[...]   Character set
[^...]  Negated set
(a|b)   Group/Alternation
`

func (m Model) ApplyRegexHighlighting(content, pattern string) string {
	if m.CompiledRegex == nil {
		return content
	}

	re := m.CompiledRegex
	lines := strings.Split(content, "\n")

	var styledLines []string
	lineOffset := 0

	for _, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		styledLine := m.ReconstructLine(line, matches, lineOffset)
		styledLines = append(styledLines, styledLine)
		lineOffset += len(line) + 1
	}

	return strings.Join(styledLines, "\n")
}

func (m Model) ReconstructLine(rawLine string, matches [][]int, lineOffset int) string {
	var b strings.Builder
	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			b.WriteString(rawLine[lastIndex:start])
		}

		matchedText := rawLine[start:end]

		currentStyle := HighlightStyle
		if len(m.Matches) > 0 {
			globalStart := lineOffset + start
			if globalStart == m.Matches[m.CurrentMatch][0] {
				currentStyle = SelectedMatchStyle
			}
		}

		b.WriteString(currentStyle.Render(matchedText))

		lastIndex = end
	}

	if lastIndex < len(rawLine) {
		b.WriteString(rawLine[lastIndex:])
	}

	return b.String()
}

func (m Model) View() string {
	helpText := helpStyle.Render("ctrl+c: quit | ↑↓/j/k: scroll | n/N: next/prev match | esc: unfocus | space: help")

	var inputView string
	if m.TextInput.Focused() {
		inputView = focusedStyle.Render(m.TextInput.View())
	} else {
		inputView = blurredStyle.Render(m.TextInput.View())
	}

	var counter string
	if len(m.Matches) > 0 {
		counter = counterStyle.Render(fmt.Sprintf("%d/%d", m.CurrentMatch+1, len(m.Matches)))
	} else if m.TextInput.Value() != "" && m.Err == nil {
		counter = counterStyle.Render("0/0")
	}

	var errorView string
	if m.Err != nil {
		errorView = errorStyle.Render(m.Err.Error())
	}

	header := inputView
	if errorView != "" {
		header = lipgloss.JoinVertical(lipgloss.Left, header, errorView)
	}

	pattern := m.TextInput.Value()
	highlightedContent := m.ApplyRegexHighlighting(m.Content, pattern)
	m.Viewport.SetContent(highlightedContent)

	mainWidth := m.Viewport.Width

	gapWidth := mainWidth - lipgloss.Width(helpText) - lipgloss.Width(counter)
	if gapWidth < 0 {
		gapWidth = 0
	}
	footer := lipgloss.JoinHorizontal(lipgloss.Top, helpText, strings.Repeat(" ", gapWidth), counter)

	ui := strings.Join([]string{
		header,
		separatorStyle.Render(strings.Repeat("—", mainWidth)),
		m.Viewport.View(),
		footer,
	}, "\n")

	if m.ShowHelp {
		helpView := sideBarStyles.Height(m.WindowHeight - 2).Render(helpContent)
		return lipgloss.JoinHorizontal(lipgloss.Top, ui, helpView)
	}

	return ui
}
