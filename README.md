# Regex Highlighter TUI

A basic terminal-based regex testing tool with real-time syntax highlighting. Test your regular expressions against any text file with instant visual feedback.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)

## Features

- 🎨 **Real-time Highlighting**: See regex matches instantly as you type
- ⚡ **Live Validation**: Immediate feedback on regex syntax errors
- 🔢 **Match Counter**: Track total matches and current position
- 🧭 **Match Navigation**: Jump between matches with `n`/`p`
- 📘 **Cheat Sheet**: Built-in regex reference guide (Press Space)

## Installation

```bash
# Install via Go
go install github.com/red7-c/regex-highlighter@latest
```

Or from source

```
# Clone the repository
git clone https://github.com/red7-c/regex-highlighter.git
cd regex-highlighter

# Build the project
go build -o regex-highlighter

# Or run directly
go run main.go <file-path>
```

## Usage

```bash
# Basic usage
./regex-highlighter path/to/your/file.txt

# Example
./regex-highlighter sample.log
```

Once running:

1. Type your regex pattern in the input field
2. Matches will be highlighted in yellow in real-time
3. Use keyboard shortcuts to navigate

## Keyboard Shortcuts

| Key             | Action                                    |
| --------------- | ----------------------------------------- |
| `Type anything` | Focus input and start typing regex        |
| `Esc`           | Unfocus input field                       |
| `↑` / `k`       | Scroll up                                 |
| `↓` / `j`       | Scroll down                               |
| `n`             | Jump to next match                        |
| `p`             | Jump to previous match                    |
| `Space`         | Toggle regex cheat sheet (when unfocused) |
| `Ctrl+C`        | Quit application                          |

## How It Works

The application reads your target file and displays it in a scrollable viewport. As you type a regex pattern:

1. **Pattern Validation**: The regex is compiled in real-time
2. **Match Finding**: All matches are located in the content
3. **Visual Highlighting**: Matches are highlighted in orange, with the selected match in blue
4. **Error Handling**: Invalid regex patterns display an error message below the input

## Project Structure

```
.
├── main.go                 # Entry point and file handling
└── internal/
    └── tui/
        ├── model.go       # Data model and initialization
        ├── update.go      # Event handling and state updates
        ├── view.go        # Rendering and highlighting logic
        └── styles.go      # UI styles
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## Examples

### Test Email Patterns

```bash
./regex-highlighter contacts.txt
# Then type: \b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b
```

### Find URLs

```bash
./regex-highlighter webpage.html
# Then type: https?://[^\s]+
```

### Match Phone Numbers

```bash
./regex-highlighter phonebook.txt
# Then type: \d{3}-\d{3}-\d{4}
```

---

## Coming Features

- [x] **Error Display** - Show regex compilation errors in the UI
- [x] **Match Counter** - Display total number of matches found
- [x] **Help Overlay** - Press `Space` to see regex cheat sheet
- [ ] **Capture Groups** - Highlight and display capture group contents
- [x] **Match Navigation** - Jump between matches with `n`/`p` keys
- [ ] **Regex Flags Support** - Toggle case-insensitive, multiline, dotall modes
- [ ] **Search History** - Navigate through previous regex patterns with arrow keys
- [ ] **Copy to Clipboard** - Copy regex or matched text with `y`/`Y`
- [ ] **Replace Mode** - Preview text replacement operations
- [ ] **Export Matches** - Save all matches to a file (JSON/CSV/TXT)
- [ ] **Multi-file Support** - Work with directories and tab between files
- [ ] **Regex Cheat Sheet** - Built-in reference for common patterns
- [ ] **Custom Color Schemes** - User-configurable highlighting colors
- [ ] **Configuration File** - Save preferences and default settings
- [ ] **Performance Optimization** - Caching for large files
- [ ] **Line Numbers** - Display line numbers in viewport
- [ ] **Context Lines** - Show N lines before/after matches
- [ ] **Regex Debugger** - Step through regex matching process
- [ ] **Pattern Library** - Save and load common regex patterns
- [ ] **Diff Mode** - Compare before/after replacement
- [ ] **Multi-pattern Support** - Test multiple regexes simultaneously
- [ ] **Syntax Highlighting** - Language-aware file display

## License

MIT License - see LICENSE file for details
