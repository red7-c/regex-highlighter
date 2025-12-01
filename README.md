# Regex Highlighter TUI

A basic terminal-based regex testing tool with real-time syntax highlighting. Test your regular expressions against any text file with instant visual feedback.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)

## Features

- 🎨 **Real-time Highlighting**: See regex matches instantly as you type
- ⚡ **Live Validation**: Immediate feedback on regex syntax errors


## Installation

```bash
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

| Key | Action |
|-----|--------|
| `Type anything` | Focus input and start typing regex |
| `Esc` | Unfocus input field |
| `↑` / `k` | Scroll up |
| `↓` / `j` | Scroll down |
| `Ctrl+C` / `q` | Quit application |

## How It Works

The application reads your target file and displays it in a scrollable viewport. As you type a regex pattern:

1. **Pattern Validation**: The regex is compiled in real-time
2. **Match Finding**: All matches are located in the content
3. **Visual Highlighting**: Matches are highlighted with a yellow background
4. **Error Handling**: Invalid regex patterns are caught (though not currently displayed)

## Project Structure

```
.
├── main.go                 # Entry point and file handling
└── internal/
    └── tui/
        ├── model.go       # Data model and initialization
        ├── update.go      # Event handling and state updates
        └── view.go        # Rendering and highlighting logic
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

### 🚀 Planned Enhancements

#### High Priority
- [ ] **Error Display** - Show regex compilation errors in the UI
- [ ] **Match Counter** - Display total number of matches found
- [ ] **Help Overlay** - Press `?` to see all keybindings
- [ ] **Capture Groups** - Highlight and display capture group contents
- [ ] **Match Navigation** - Jump between matches with `n`/`N` keys

#### Medium Priority
- [ ] **Regex Flags Support** - Toggle case-insensitive, multiline, dotall modes
- [ ] **Search History** - Navigate through previous regex patterns with arrow keys
- [ ] **Copy to Clipboard** - Copy regex or matched text with `y`/`Y`
- [ ] **Replace Mode** - Preview text replacement operations
- [ ] **Export Matches** - Save all matches to a file (JSON/CSV/TXT)

#### Nice to Have
- [ ] **Multi-file Support** - Work with directories and tab between files
- [ ] **Regex Cheat Sheet** - Built-in reference for common patterns
- [ ] **Custom Color Schemes** - User-configurable highlighting colors
- [ ] **Configuration File** - Save preferences and default settings
- [ ] **Performance Optimization** - Caching for large files
- [ ] **Line Numbers** - Display line numbers in viewport
- [ ] **Context Lines** - Show N lines before/after matches
- [ ] **Benchmark Mode** - Measure regex performance on large files

#### Advanced Features
- [ ] **Regex Debugger** - Step through regex matching process
- [ ] **Pattern Library** - Save and load common regex patterns
- [ ] **Diff Mode** - Compare before/after replacement
- [ ] **Multi-pattern Support** - Test multiple regexes simultaneously
- [ ] **Syntax Highlighting** - Language-aware file display
- [ ] **Remote Files** - Open files via HTTP/HTTPS URLs

## Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Priorities

If you'd like to contribute, consider tackling these first:
1. Fix error display in UI (see `view.go`)
2. Add match counter functionality
3. Implement help overlay
4. Add unit tests

## Known Issues

- Regex errors are validated but not displayed to users
- Character limit is set to 156 (may be too restrictive)
- No support for regex flags (case-insensitive, etc.)
- Large files may cause performance issues

## License

MIT License - see LICENSE file for details

## Acknowledgments

Built with the excellent [Charm](https://charm.sh/) TUI libraries.

## Contact

- GitHub: [@red7-c](https://github.com/red7-c)
- Project Link: [https://github.com/red7-c/regex-highlighter](https://github.com/red7-c/regex-highlighter)

---

**Star ⭐ this repository if you find it useful!**