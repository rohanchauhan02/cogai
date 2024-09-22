package create

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rohanchauhan02/cogai/internals/pkg/ui/option"
	"github.com/spf13/cobra"
)

// Original Logo text
var Logo = `
   ____    _____     ______       ____      ___
  / ___|  / ___ \   / _____|     / /\ \     | |
 | |     | |   | | | |    _     / /__\ \    | |
 | |___  | |___| | | |___| |   / /----\ \   | |
  \____|  \_____/   \____|_|  |_|      |_|  |_|

                          ...COGNATIVE AI ENGNE


                Version 0.0.1
       Copyright (c) 2024 Rohan Chauhan
        <singhrohankumar7@gmail.com>

      CogAI: AI-Driven Command-Line Tool.
`

// Define red and yellow styles using lipgloss
var (
	RedStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000")) // Red color
	YellowStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFF00")) // Yellow color
)

// CreateCmd is the main command for initializing CogAI
var CreateCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial setup for CogAI",
	Long:  `Initial setup for CogAI`,
	Run: func(cmd *cobra.Command, args []string) {
		BlinkLogo() // Blink the logo in red and yellow
		p := tea.NewProgram(option.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func BlinkLogo() {
	for i := 0; i < 10; i++ { // Loop to create blink effect
		clearScreen() // Clear the screen each time
		if i%2 == 0 {
			fmt.Println(RedStyle.Render(Logo)) // Display the logo in red
		} else {
			fmt.Println(YellowStyle.Render(Logo)) // Display the logo in yellow
		}
		time.Sleep(100 * time.Millisecond) // Pause for half a second
	}
	clearScreen()
	fmt.Println(YellowStyle.Render(Logo)) // Final display in red after blinking ends
}
