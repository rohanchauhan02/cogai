package aioutput

// import (
// 	"fmt"

// 	tea "github.com/charmbracelet/bubbletea"
// 	aitext "github.com/rohanchauhan02/cogai/internals/pkg/ai/text"
// )

// // OutputModel is the model that displays output based on user input
// type OutputModel struct {
// 	value string
// }

// // TextOutputModel initializes the output model with the user input
// func TextOutputModel(value string) OutputModel {
// 	return OutputModel{
// 		value: value,
// 	}
// }

// // Init initializes the tea program for the output model
// func (m OutputModel) Init() tea.Cmd {
// 	return nil
// }

// // Update handles the output and returns back to input model after displaying output
// func (m OutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "enter", " ":
// 			// Go back to input model to ask for the next input
// 			fmt.Println(m.value)
// 			return aitext.TextAIInputModel(m.value), nil
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		}
// 	}
// 	return m, nil
// }

// // View displays the output and waits for user to press "Enter" to continue
// func (m OutputModel) View() string {
// 	s := fmt.Sprintf("You entered: %s\n", m.value)
// 	s += "\nPress Enter to input again or ctrl+c to quit.\n"
// 	return s
// }
