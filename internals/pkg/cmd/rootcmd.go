package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type ICmd interface {
	Execute()
	AddCommand(*cobra.Command)
}

type cmd struct {
	rootCmd *cobra.Command
}

func NewCmd() ICmd {
	return &cmd{
		rootCmd: &cobra.Command{
			Use:   "cogai",
			Short: "CogAI is an AI-enabled CLI tool that interacts with OpenAI to answer your questions",
			Long:  `CogAI is a GoLang-based AI CLI that allows you to interact with OpenAI's for smart and insightful responses.`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					// provide the version and all about the CLI tool
					fmt.Println(`
						CogAI: AI-Enabled CLI Tool

						CogAI is an advanced command-line interface (CLI) tool powered by OpenAI, enabling intelligent interactions to answer your queries with insightful responses. Built with GoLang, CogAI streamlines communication with AI for efficient, real-time solutions.

						Version: 0.0.1
						Author: Rohan Chauhan
						GitHub: https://github.com/rohanchauhna02/cogai
						License: MIT

						Usage:
							cogai [command] [options]

						Commands:
							ask        - Submit a question for AI to answer
							ask <question>  - Ask a specific question
							help       - Display available commands and usage information
							help <command>  - Get detailed help for a specific command
							version    - Display the current version of CogAI
							version <command> - Show version information for a specific command

						Examples:
							cogai ask "What is the meaning of life?"
					`)

					return
				}
			},
		},
	}
}

func (c *cmd) AddCommand(cmd *cobra.Command) {
	c.rootCmd.AddCommand(cmd)
}

func (c *cmd) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
