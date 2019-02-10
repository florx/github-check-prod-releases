package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

var client *github.Client
var ctx context.Context

func main() {

	_, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		panic("The ENV variable GITHUB_TOKEN is not set.")
	}

	_, exists = os.LookupEnv("GITHUB_ORG")
	if !exists {
		panic("The ENV variable GITHUB_ORG is not set.")
	}

	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

	var allRepos []*github.Repository

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, os.Getenv("GITHUB_ORG"), opt)
		if err != nil {
			println(err.Error())
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, repo := range allRepos {

		//skip archived
		if !repo.GetArchived() {
			processRepo(repo)
		}
	}

}

func processRepo(repo *github.Repository) {

	comparison, response, err := client.Repositories.CompareCommits(ctx, repo.GetOwner().GetLogin(), repo.GetName(), "master", "prod")
	if err != nil {
		if response.StatusCode != 404 {
			println(err.Error())
			return
		}
	}

	if comparison.GetBehindBy() > 0 {
		fmt.Printf("%s is behind by %d releases\n", repo.GetName(), comparison.GetBehindBy())
	}
}
