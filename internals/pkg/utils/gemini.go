package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rohanchauhan02/cogai/internals/modules/env"
)

type GeminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content string `json:"content"`
	} `json:"candidates"`
	Error GeminiError `json:"error,omitempty"`
}

func AskGemini(prompt string) (string, error) {
	url := "https://generativelanguage.googleapis.com/v1beta2/models/text-bison-001:generateText"

	// Prepare the request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt": map[string]string{
			"text": prompt,
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

	apiKey := env.GetKey("GEMINI_API_KEY", true)
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

	var result GeminiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Handle potential error in response
	if result.Error.Message != "" {
		return "", errors.New(result.Error.Message)
	}

	// Extract the message content from the candidates
	if len(result.Candidates) > 0 {
		message := result.Candidates[0].Content
		return message, nil
	}

	return "", errors.New("no response from Gemini")
}
