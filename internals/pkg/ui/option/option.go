package option

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rohanchauhan02/cogai/internals/pkg/ui/text"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		choices:  []string{"ChatGPT-4", "Google Gemini"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	fmt.Println("Initializing model...")
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			// Select the model and proceed to get the API key
			modelSelected := m.choices[m.cursor]
			fmt.Printf("You selected: %s\n", modelSelected)
			return text.TextInputModel(modelSelected), nil
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which AI model would you like to use?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress 'q' to exit and 'Enter' to select a model.\n"
	return s
}
