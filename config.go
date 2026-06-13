package main

import (
	"log"
	"os"
)

type Config struct {
	GitHubToken string
	GroqKey     string
}

func loadConfig() Config {
	gh := os.Getenv("GITHUB_TOKEN")
	gk := os.Getenv("GROQ_API_KEY")

	if gh == "" || gk == "" {
		log.Fatal("GITHUB_TOKEN and GROQ_API_KEY must be set")
	}

	return Config{GitHubToken: gh, GroqKey: gk}
}
