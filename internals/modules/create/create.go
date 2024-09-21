package create

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rohanchauhan02/cogai/internals/pkg/ui/option"
	"github.com/spf13/cobra"
)

var Logo = `
   ____    _____     ______       ____      ___
  / ___|  / ___ \   / _____|     / /\ \     | |
 | |     | |   | | | |    _     / /__\ \    | |
 | |___  | |___| | | |___| |   / /----\ \   | |
  \____|  \_____/   \____|_|  |_|      |_|  |_|


                Version 0.0.1
       Copyright (c) 2024 Rohan Chauhan
        <singhrohankumar7@gmail.com>

      CogAI: AI-Driven Command-Line Tool.
`

var (
	infoStyle = lipgloss.NewStyle().Bold(true).Border(lipgloss.HiddenBorder()).Foreground(lipgloss.Color("#fc9e21"))
)

var CreateCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial setup for CogAI",
	Long:  `Initial setup for CogAI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(infoStyle.Render(Logo))
		p := tea.NewProgram(option.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}
