package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v29/github"
	"github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Fatal("hit rate limit")
	}
	if _, ok := err.(*github.AcceptedError); ok {
		log.Fatal("scheduled on GitHub side")
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range repos {
		pp.Println(v)
	}
}
