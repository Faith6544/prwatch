package main

import (
	"fmt"
	"time"
)

func watch(repo string, cfg Config) {
	seen := map[int64]bool{}
	fmt.Printf("👀 watching PRs on %s...\n\n", repo)

	for {
		prs, err := getOpenPRs(repo, cfg.GitHubToken)
		if err != nil {
			fmt.Println("error fetching PRs:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		if len(prs) == 0 {
			fmt.Println("no open PRs found, checking again in 30s...")
			time.Sleep(30 * time.Second)
			continue
		}

		for _, pr := range prs {
			comments, err := getReviewComments(repo, pr.Number, cfg.GitHubToken)
			if err != nil {
				continue
			}

			for _, c := range comments {
				if seen[c.ID] {
					continue
				}
				seen[c.ID] = true

				fmt.Printf("─────────────────────────────────\n")
				fmt.Printf("📌 PR #%d: %s\n", pr.Number, pr.Title)
				fmt.Printf("💬 @%s commented on %s:\n   %s\n\n", c.User.Login, c.Path, c.Body)
				fmt.Printf("🤖 suggested reply:\n")

				reply, err := suggestReply(c, cfg.GroqKey)
				if err != nil {
					fmt.Println("   (couldn't generate reply:", err, ")")
				} else {
					fmt.Println("   ", reply)
				}
				fmt.Println()
			}
		}

		time.Sleep(30 * time.Second)
	}
}
