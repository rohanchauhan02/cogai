package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (

	// Cobra command to get environment variables
	GetEnv = &cobra.Command{
		Use:   "get",
		Short: "Retrieve environment variables from the configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			// If the --all flag is set, retrieve all keys
			showAll, _ := cmd.Flags().GetBool("all")
			if showAll {
				getAllKeys()
				return
			}

			// If no key is provided via argument or flag, show help
			if len(args) == 0 && key == "" {
				cmd.Help()
				return
			}

			// Use the key from the argument if provided, otherwise use the flag
			if key == "" && len(args) > 0 {
				key = strings.ToUpper(args[0])
			}

			// Retrieve the specific key from the configuration
			GetKey(key, false)
		},
	}
)

// Retrieve a specific key from the config file
func GetKey(key string, hide bool) string {
	// Check if config file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Error: config file not found")
		return ""
	}

	// Read the config file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return ""
	}

	// Unmarshal the YAML data
	var config map[string]string
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return ""
	}

	// Check if the requested key exists
	value, ok := config[key]
	if !ok {
		fmt.Println("Error: key not found")
		return ""
	}

	// Print the key's value
	if !hide {
		fmt.Printf("%s: %s\n", key, value)
	}
	return value
}

// Retrieve all keys from the config file
func getAllKeys() {
	// Check if config file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Error: config file not found")
		return
	}

	// Read the config file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Unmarshal the YAML data
	var config map[string]string
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}

	// Print all keys and values
	for key, value := range config {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func init() {
	// Flag to retrieve all keys
	GetEnv.Flags().BoolP("all", "a", false, "Retrieve all environment variables")

	// Flag to specify a key to retrieve
	GetEnv.Flags().StringVarP(&key, "key", "k", "", "Specify key to retrieve")
}
