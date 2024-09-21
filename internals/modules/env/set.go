package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	ExportCmd = &cobra.Command{
		Use:   "set",
		Short: "set the environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			arguments := strings.Split(args[0], "=")
			saveAPIKey(arguments[0], arguments[1])
		},
	}
	filePath = "app.config.yml"
)

func saveAPIKey(key, value string) {
	config := make(map[string]string)
	key = strings.ToUpper(key)
	DeleteEnvVar(key, false)
	config[key] = value
	data, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error marshaling YAML:", err)
		return
	}
	// create if not exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	}
	// Permissions for security
	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		fmt.Println("Error writing key:", err)
		return
	}
	fmt.Println("key saved successfully.")
}
