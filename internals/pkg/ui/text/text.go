package text

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rohanchauhan02/cogai/internals/modules/env"
)

var (
	placeHolderStyle = lipgloss.NewStyle().Italic(true)
)

type textModel struct {
	textInput textinput.Model
	modelName string
}

func TextInputModel(modelName string) textModel {
	ti := textinput.New()
	ti.Placeholder = "Enter the API key"
	ti.PlaceholderStyle = placeHolderStyle
	ti.Focus()
	return textModel{
		textInput: ti,
		modelName: modelName,
	}
}

func (m textModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			value := strings.Split(m.textInput.Value(), " ")[0]

			// Store the API key or perform actions here
			key := "OPENAI_API_KEY"
			if m.modelName == "ChatGPT-4o" {
				key = "OPENAI_API_KEY"
			} else if m.modelName == "Google Gemini" {
				key = "GOOGLE_API_KEY"
			}
			env.SaveAPIKey(key, value)
			return m, tea.Quit
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textModel) View() string {
	s := "Please Enter the API key?\n\n"
	s += m.textInput.View() + "\n"
	s += "\nPress ctrl+c to exit.\n"
	return s
}
