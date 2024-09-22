package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rohanchauhan02/cogai/internals/modules/env"
)

func AskChatGPT(prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// Prepare the request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	apiKey := env.GetKey("OPENAI_API_KEY", true)
	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and process the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Extract the message content
	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return message, nil
}
