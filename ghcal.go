package ghcal

import (
	"fmt"
	"log"
)

// Debug is debug flag
var Debug bool

func init() {
	Debug = false
}

func getCommits(org string, repo *string) ([]*Commit, error) {
	var repos []string
	if repo == nil {
		var err error
		repos, err = getGithubOrgRepos(org)
		if err != nil {
			return nil, err
		}
	} else {
		repos = []string{*repo}
	}
	allCommits := []*Commit{}
	for _, repo := range repos {
		commits, err := getGithubCommits(org, repo)
		if err != nil {
			return nil, err
		}
		allCommits = append(allCommits, commits...)
	}
	return allCommits, nil
}

func getCalendar(commits []*Commit) string {
	cal := Calendar{
		Prodid: "-//github.com/yushi//ghcal//EN",
	}
	for _, c := range commits {
		ev := Event{
			UID:         fmt.Sprintf("%s-%s", c.RepoName, c.Hash),
			Summary:     &c.RepoName,
			Description: &c.Message,
			DTStamp:     c.CommitedAt,
			DTStart:     c.CommitedAt,
		}
		cal.Events = append(cal.Events, ev)
	}
	return cal.ICalString()
}

// Calendarize returns calendar
func Calendarize(org string, repo *string) {
	commits, err := getCommits(org, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getCalendar(commits))
}
