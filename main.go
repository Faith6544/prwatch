package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	repo := flag.String("repo", "", "GitHub repo in owner/repo format")
	flag.Parse()

	if *repo == "" {
		fmt.Println("usage: prwatch --repo owner/repo")
		os.Exit(1)
	}

	cfg := loadConfig()
	watch(*repo, cfg)
}
