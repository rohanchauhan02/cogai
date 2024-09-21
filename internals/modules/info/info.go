package info

import "github.com/spf13/cobra"

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about CogAI",
	Long:  `Get information about CogAI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
