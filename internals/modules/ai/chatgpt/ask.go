package chatgpt

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rohanchauhan02/cogai/internals/modules/create"
	aitext "github.com/rohanchauhan02/cogai/internals/pkg/ai/text"

	"github.com/spf13/cobra"
)

var (
	AskCmd = &cobra.Command{
		Use:   "ask",
		Short: "Ask a question to the AI",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println(create.InfoStyle.Render(create.Logo))
			create.BlinkLogo()
			// if len(args) == 0 {
			// 	cmd.Help()
			// 	return
			// }

			p := tea.NewProgram(aitext.TextAIInputModel("TextAIModel"))
			if _, err := p.Run(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// question := strings.Join(args, " ")
			// response, err := askChatGPT(question)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// 	return
			// }
			// fmt.Println(response)
		},
	}
)
