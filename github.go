package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Comment struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
	Path string `json:"path"`
	Line int    `json:"line"`
	User struct {
		Login string `json:"login"`
	} `json:"user"`
	DiffHunk string `json:"diff_hunk"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
}

func githubGet(url, token string, out any) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return json.Unmarshal(body, out)
}

func getOpenPRs(repo, token string) ([]PullRequest, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls?state=open", repo)
	var prs []PullRequest
	err := githubGet(url, token, &prs)
	return prs, err
}

func getReviewComments(repo string, prNumber int, token string) ([]Comment, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d/comments", repo, prNumber)
	var comments []Comment
	err := githubGet(url, token, &comments)
	return comments, err
}
