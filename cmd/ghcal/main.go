package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yushi/ghcal"
)

type conf struct {
	Token string
}

func main() {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/.ghcal", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}
	c := conf{}
	if err := json.Unmarshal(b, &c); err != nil {
		log.Fatal(err)
	}
	ghcal.GithubToken = c.Token

	if len(os.Args) < 2 {
		log.Fatal("org or org/repo required.")
	}
	var (
		org  string
		repo *string
	)
	input := os.Args[1]

	if idx := strings.Index(input, "/"); idx != -1 {
		org = input[0:idx]
		r := input[idx+1:]
		repo = &r
	} else {
		org = input
	}
	ghcal.Calendarize(org, repo)
}
