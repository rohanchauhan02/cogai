package chatgpt

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/rohanchauhan02/cogai/internals/modules/create"
	aitext "github.com/rohanchauhan02/cogai/internals/pkg/ai/text"

	"github.com/rohanchauhan02/cogai/internals/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	AskCmd = &cobra.Command{
		Use:   "ask [flags] [question]",
		Short: "Ask a question to the AI",
		Long: `Use this command to interact with the AI.

  For interactive mode:
    cogai ask -i

  For direct question mode:
    cogai ask "Your question here"`,
		Run: func(cmd *cobra.Command, args []string) {

			interactive, _ := cmd.Flags().GetBool("interactive")

			if interactive {
				// Start the interactive mode
				create.BlinkLogo()
				p := tea.NewProgram(aitext.TextAIInputModel("TextAIModel"))
				if _, err := p.Run(); err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			} else if len(args) > 0 {
				// Process the direct question
				question := strings.Join(args, " ")
				response, err := utils.AskChatGPT(question)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Println(response)
			} else {
				cmd.Help()
				return
			}
		},
	}
)

func init() {
	AskCmd.Flags().BoolP("interactive", "i", false, "Start interactive mode")
}
