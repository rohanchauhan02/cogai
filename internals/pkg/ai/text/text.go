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
	borderStyle      = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1).Margin(1)
	boxStyle         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(1)
	exitInstructions = "\nPress ctrl+c to exit, ⌃ for scroll up and ⌄ for scroll down."

	// Set the maximum width for the response box
	maxWidth = 80
)

const responseHeight = 100 // Number of lines to show in the response box

type TextAIModel struct {
	textInput     textinput.Model
	modelName     string
	response      string
	scrollOffset  int
	responseLines []string
}

// TextAIInputModel initializes the text input model
func TextAIInputModel(modelName string) TextAIModel {
	ti := textinput.New()
	ti.Placeholder = "Ask a question or enter code..."
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
		switch msg.Type {
		case tea.KeyEnter:
			value := strings.TrimSpace(m.textInput.Value())
			if value != "" {
				// Get the response
				response, err := utils.AskChatGPT(value)
				if err != nil {
					fmt.Println("Error:", err)
					return m, tea.Quit
				}

				// Split the response into lines for scrolling and wrapping
				m.responseLines = wrapText(response, maxWidth)
				m.scrollOffset = 0 // Reset scroll offset
				m.response = response
				m.textInput.SetValue("") // Clear input field
				return m, nil
			}
		case tea.KeyUp:
			if m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case tea.KeyDown:
			if m.scrollOffset < len(m.responseLines)-responseHeight {
				m.scrollOffset++
			}
		}
		switch msg.String() {
		case "ctrl+c", "q":
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

		responseContent := strings.Join(m.responseLines[start:end], "\n")
		output += borderStyle.Render(responseStyle.Render("AI Response:")+"\n"+responseContent) + "\n"
	}

	// Box for the input prompt
	inputBox := boxStyle.Render(inputStyle.Render("Please ask a question or enter code:\n") + m.textInput.View())

	// Combine the response and input box
	output += inputBox + "\n"

	// Display exit instructions
	output += exitInstructions

	return output
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
