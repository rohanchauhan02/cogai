package ai

import (
	"github.com/spf13/cobra"
)

var (
	AskAiCmd = &cobra.Command{
		Use:   "ask",
		Short: "Ask a question to the AI",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			// question := strings.Join(args, " ")
			// ai.Ask(question)
		},
	}
)
