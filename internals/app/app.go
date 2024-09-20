package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Init() {

	// Create a new Cobra CLI app
	var rootCmd = &cobra.Command{
		Use:   "cogai",
		Short: "CogAI is an AI-enabled CLI tool that interacts with OpenAI to answer your questions",
		Long:  `CogAI is a GoLang-based AI CLI that allows you to interact with OpenAI's ChatGPT for smart and insightful responses.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a prompt to ask the AI")
				return
			}
			fmt.Printf("Hello %s!, I'm CogAI, your AI-powered CLI tool!", args[0])
		},
	}

	// Execute the CLI tool
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
