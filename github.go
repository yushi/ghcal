package ghcal

import (
	"log"
	"strings"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

type Commit struct {
	Hash       string
	RepoName   string
	CommitedAt time.Time
	Message    string
	Committer  string
}

var GithubToken string

func getGithubClient() *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: GithubToken},
	}
	return github.NewClient(t.Client())
}
func getGithubOrgRepos(org string) ([]string, error) {
	client := getGithubClient()
	//opt := &github.RepositoryListByOrgOptions{Sort: "updated"}
	repos, resp, err := client.Repositories.ListByOrg(org, nil)
	if err != nil {
		return nil, err
	}
	if Debug {
		log.Print(resp.Rate)
	}
	repoNames := make([]string, len(repos))
	for i, r := range repos {
		repoNames[i] = *r.Name
	}
	return repoNames, nil
}

func getGithubCommits(org, repo string) ([]*Commit, error) {
	client := getGithubClient()
	repoCommits, resp, err := client.Repositories.ListCommits(org, repo, nil)
	if err != nil {
		return nil, err
	}
	if Debug {
		log.Print(resp.Rate)
	}
	commits := make([]*Commit, len(repoCommits))
	for i, c := range repoCommits {
		commit := &Commit{
			Hash:     *c.SHA,
			RepoName: org + "/" + repo,
			Message:  strings.Split(*c.Commit.Message, "\n")[0],
		}
		if c.Committer != nil {
			if c.Committer.Email != nil {
				commit.Committer = *c.Committer.Email
			} else if c.Committer.Name != nil {
				commit.Committer = *c.Committer.Name
			}
		}
		commit.CommitedAt = *c.Commit.Committer.Date

		commits[i] = commit
	}
	return commits, nil
}
