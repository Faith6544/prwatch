# prwatch

A CLI agent that watches your GitHub pull requests and suggests replies to review comments using AI.

## What it does

- Polls your open PRs every 30 seconds
- Detects new review comments as they come in
- Feeds the comment + diff context to an LLM
- Prints a suggested reply right in your terminal

## Setup

1. Get a GitHub personal access token from github.com/settings/tokens (needs `repo` scope)
2. Get a free Groq API key from console.groq.com
3. Set your environment variables:

**Mac/Linux:**
```bash
export GITHUB_TOKEN=your_github_token
export GROQ_API_KEY=your_groq_key
```

**Windows (PowerShell):**
```powershell
$env:GITHUB_TOKEN = "your_github_token"
$env:GROQ_API_KEY = "your_groq_key"
```

## Run

```bash
go run . --repo owner/repo-name
```

## Build

```bash
go build -o prwatch .
./prwatch --repo owner/repo-name
```

## Stack

- Go
- GitHub REST API
- Groq API (llama-3.3-70b-versatile)