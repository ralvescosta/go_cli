package app

import (
	"fmt"
	"go_cli/pkg/infra/repo"
	"log"
)

type GithunStatusApp struct {
	repo repo.GetGithubUserDetails
}

func (g GithunStatusApp) ShowGithubUserDetails(user string) {
	result, err := g.repo.Get(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bio: " + result.Bio)
	fmt.Println("Followers: " + fmt.Sprint(result.Followers))
	fmt.Println("Following: " + fmt.Sprint(result.Following))
}

func NewGithubStatusApp(repo repo.GetGithubUserDetails) GithunStatusApp {
	return GithunStatusApp{repo: repo}
}
