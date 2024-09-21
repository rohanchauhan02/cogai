package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	key       string
	DeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a specific environment variable",
		Run: func(cmd *cobra.Command, args []string) {
			if key == "" {
				fmt.Println("Error: key is required")
				return
			}
			// Call the function to delete the key from the config
			DeleteEnvVar(strings.ToUpper(key), true)
		},
	}
)

// delete the env variable from the config file
func DeleteEnvVar(key string, isExist bool) {
	// check if the key exists in the config file
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Error: config file not found")
		return
	}
	// read the config file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	// unmarshal the config
	var config map[string]string
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
	// check if the key exists in the config
	if _, ok := config[key]; !ok {
		if isExist {
			fmt.Println("Error: key not found")
		}
		return
	}
	// delete the
	delete(config, key)
	// marshal the config
	data, err = yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error marshaling config:", err)
		return
	}
	// write the config
	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		fmt.Println("Error writing config:", err)
		return
	}
	if isExist {
		fmt.Println("Key deleted successfully.")
	}
}
func init() {
	DeleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key to delete")
	DeleteCmd.MarkFlagRequired("key")
}
