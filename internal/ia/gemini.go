package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ContentPart struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []ContentPart `json:"parts"`
}

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

func CallGemini(prompt string) (string, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

	reqBody := GeminiRequest{
		Contents: []Content{
			{Parts: []ContentPart{{Text: prompt}}},
		},
	}

	jsonData, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", "AIzaSyB2-NQSkKiObstbSbiFfR3FcMmIperKZ84")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Marshal de volta para JSON e transforma em string
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
