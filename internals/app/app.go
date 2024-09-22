package app

import (
	"github.com/rohanchauhan02/cogai/internals/modules/ai/chatgpt"
	"github.com/rohanchauhan02/cogai/internals/modules/create"
	"github.com/rohanchauhan02/cogai/internals/modules/env"
	"github.com/rohanchauhan02/cogai/internals/modules/info"
	"github.com/rohanchauhan02/cogai/internals/pkg/cmd"
)

func Init() {
	rootCmd := cmd.NewCmd()
	rootCmd.AddCommand(chatgpt.AskCmd)
	rootCmd.AddCommand(chatgpt.AskCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(env.ExportCmd)
	rootCmd.AddCommand(env.DeleteCmd)
	rootCmd.AddCommand(env.GetEnv)
	rootCmd.Execute()
}
