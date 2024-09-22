package aitext

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rohanchauhan02/cogai/internals/pkg/utils"
)

var (
	placeHolderStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("240"))
	inputStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF"))
	responseStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#ADFF2F"))
	codeStyle        = lipgloss.NewStyle().Background(lipgloss.Color("#1E1E1E")).Foreground(lipgloss.Color("#C5C5C5")).Padding(1)
	borderStyle      = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1).Margin(1)
	boxStyle         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(1)
	exitInstructions = "\nPress ctrl+c to exit, ⌃ for scroll up and ⌄ for scroll down."

	// Set the maximum width for the response box
	maxWidth = 100
)

const responseHeight = 10 // Number of lines to show in the response box

type TextAIModel struct {
	textInput     textinput.Model
	modelName     string
	scrollOffset  int
	responseLines []string
	isMultiline   bool
}

// TextAIInputModel initializes the text input model
func TextAIInputModel(modelName string) TextAIModel {
	ti := textinput.New()
	ti.Placeholder = "Ask a anything..."
	ti.PlaceholderStyle = placeHolderStyle
	ti.Focus()
	return TextAIModel{
		textInput: ti,
		modelName: modelName,
	}
}

// Init initializes the tea program for text input
func (m TextAIModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update processes input and switches to output model after 'Enter'
func (m TextAIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// If we're in multi-line mode, allow "Enter" to insert a newline
			if m.isMultiline {
				m.textInput.SetValue(m.textInput.Value() + "\n")
			} else {
				// Handle the input submission
				value := strings.TrimSpace(m.textInput.Value())
				if value != "" {
					response, err := utils.AskChatGPT(value)
					if err != nil {
						// Clear input after error
						m.textInput.SetValue("")
						m.responseLines = wrapText(fmt.Sprintf("Error: %v", err), maxWidth)
						return m, nil
					}
					// Process successful response
					m.responseLines = wrapText(response, maxWidth)
					m.scrollOffset = 0
					m.textInput.SetValue("") // Clear input field after success
				}
			}
		case "shift+enter":
			// Toggle multi-line input mode
			m.isMultiline = true
			m.textInput.SetValue(m.textInput.Value() + "\n")
		case "up":
			// Scroll up in the response
			if m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			// Scroll down in the response
			if m.scrollOffset < len(m.responseLines)-responseHeight {
				m.scrollOffset++
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View displays the text input and response view
func (m TextAIModel) View() string {
	var output string

	// If there's a response, display it with scrolling
	if len(m.responseLines) > 0 {
		start := m.scrollOffset
		end := start + responseHeight
		if end > len(m.responseLines) {
			end = len(m.responseLines)
		}

		responseContent := formatResponse(strings.Join(m.responseLines[start:end], "\n"))
		output += borderStyle.Render(responseStyle.Render("AI Response:") + "\n" + responseContent) + "\n"
	}

	// Box for the input prompt, without extra space
	inputBox := boxStyle.Render(inputStyle.Render("Please ask a question:") + "\n" + m.textInput.View())

	// Combine the response and input box
	output += inputBox + "\n"

	// Display exit instructions
	output += exitInstructions

	return output
}

// formatResponse formats the response, identifying code blocks
func formatResponse(response string) string {
	lines := strings.Split(response, "\n")
	var formattedLines []string

	inCodeBlock := false
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock // Toggle code block state
			if inCodeBlock {
				formattedLines = append(formattedLines, codeStyle.Render("")) // Start code block
			} else {
				formattedLines = append(formattedLines, "") // End code block
			}
		} else if inCodeBlock {
			// Indent code lines for better visibility
			formattedLines = append(formattedLines, codeStyle.Render(line))
		} else {
			// Regular text line, include it without special handling
			formattedLines = append(formattedLines, line)
		}
	}

	// Ensure we add a final line if we're still in a code block
	if inCodeBlock {
		formattedLines = append(formattedLines, "") // Close the code block
	}

	return strings.Join(formattedLines, "\n")
}

// wrapText wraps the given text into lines of specified width
func wrapText(text string, width int) []string {
	var wrappedLines []string
	words := strings.Fields(text)
	var currentLine []string

	for _, word := range words {
		if lipgloss.Width(strings.Join(currentLine, " "))+lipgloss.Width(word) > width {
			wrappedLines = append(wrappedLines, strings.Join(currentLine, " "))
			currentLine = []string{word}
		} else {
			currentLine = append(currentLine, word)
		}
	}

	// Add the last line if there's any content left
	if len(currentLine) > 0 {
		wrappedLines = append(wrappedLines, strings.Join(currentLine, " "))
	}

	return wrappedLines
}
