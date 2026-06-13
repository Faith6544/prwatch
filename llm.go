package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func suggestReply(comment Comment, apiKey string) (string, error) {
	prompt := fmt.Sprintf(`You are a helpful code reviewer assistant.

A reviewer left this comment on a pull request:
"%s"

The diff context around that line:
%s

Write a short, professional reply to this review comment.
Either acknowledge the feedback and explain your approach,
or agree and mention you'll fix it. Keep it under 3 sentences.`,
		comment.Body, comment.DiffHunk)

	payload := GroqRequest{
		Model:    "llama-3.3-70b-versatile",
		Messages: []Message{{Role: "user", Content: prompt}},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result GroqResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("empty response from Groq")
	}

	return result.Choices[0].Message.Content, nil
}
