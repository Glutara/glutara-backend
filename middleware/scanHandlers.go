package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Candidate struct {
	Content       Content `json:"content"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
	FinishReason  string `json:"finishReason,omitempty"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type SafetyRating struct {
	Category         string  `json:"category"`
	Probability      string  `json:"probability"`
	ProbabilityScore float64 `json:"probabilityScore"`
	Severity         string  `json:"severity"`
	SeverityScore    float64 `json:"severityScore"`
}

func ScanFood(response http.ResponseWriter, request *http.Request) {
	// Parse JSON request body
	var requestBody map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(response, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// Extract image URL from request body
	imageURL, ok := requestBody["image_url"].(string)
	if !ok {
		http.Error(response, "Invalid image URL", http.StatusBadRequest)
		return
	}

	// Prepare request body
	bodyData := map[string]interface{}{
		"contents": map[string]interface{}{
			"role": "user",
			"parts": []map[string]interface{}{
				{
					"fileData": map[string]interface{}{
						"mimeType": "image/jpeg",
						"fileUri": imageURL,
					},
				},
				{
					"text": "Describe this picture.",
				},
			},
		},
		"safety_settings": map[string]interface{}{
			"category":  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
			"threshold": "BLOCK_LOW_AND_ABOVE",
		},
		"generation_config": map[string]interface{}{
			"temperature":      0.4,
			"topP":             1.0,
			"topK":             32,
			"maxOutputTokens":  2048,
		},
	}

	fmt.Println(bodyData)
	requestBodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		http.Error(response, "Failed to marshal request body", http.StatusInternalServerError)
		return
	}

	// Create HTTP request
	fmt.Println(requestBodyBytes)
	url := os.Getenv("GEMINI_MODEL")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		http.Error(response, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set authorization header from .env
	fmt.Println(req)
	bearerToken := os.Getenv("BEARER_TOKEN")
	if bearerToken == "" {
		http.Error(response, "BEARER_TOKEN is not set", http.StatusInternalServerError)
		return
	}

	fmt.Println(bearerToken)
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(response, "Failed to send request", http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
	fmt.Println(res.Body)
	defer res.Body.Close()

	// Read response
	var candidates []Candidate
	err = json.NewDecoder(res.Body).Decode(&candidates)
	if err != nil {
		http.Error(response, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	// Extract text values from parts
	var textValues []string
	for _, candidate := range candidates {
		for _, part := range candidate.Content.Parts {
			fmt.Println(part)
			textValues = append(textValues, part.Text)
		}
	}

	// Marshal text values to JSON
	textJSON, err := json.Marshal(textValues)
	if err != nil {
		http.Error(response, "Failed to marshal text values", http.StatusInternalServerError)
		return
	}

	// Forward text values to client
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Println(candidates)
	response.Write(textJSON)
}