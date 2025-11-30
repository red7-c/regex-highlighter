package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/red7-c/regex-highlighter/internal/tui"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run main.go <file-path>")
		os.Exit(1)
	}
	filepath := os.Args[1]
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error:\t", err)
	}
	filePathElements := strings.Split(filepath, "/")
	p := tea.NewProgram(
		tui.InitialModel(
			filePathElements[len(filePathElements)-1],
			string(content),
		),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
