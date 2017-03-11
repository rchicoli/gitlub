package github

import (
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubApi struct {
	Token string
}

func (g *GithubApi) FindRepository() ([]*github.Repository, error) {
	// list all repositories for the authenticated user
	// func (s *RepositoriesService) List(user string, opt *RepositoryListOptions) ([]*Repository, *Response, error)

	// [github.Repository{ID:42431301 ..
	// github.Rate{Limit:5030, Remaining:4296, Reset:github.Timestamp{2016-07-05 20:19:34 +0200 CEST}}

	// list all repositories for the authenticated user
	var allRepos []*github.Repository
	ch := make(chan []*github.Repository)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	// client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			PerPage: 10,
		},
	}

	go func() {
		// get all pages of results
		for {
			repos, resp, err := client.Repositories.List("rchicoli", opt)
			if err != nil {
				break
			}
			ch <- repos
			if resp.NextPage == 0 {
				close(ch)
				break
			}
			opt.ListOptions.Page = resp.NextPage
		}
	}()

	for {
		repos, ok := <-ch
		if ok {
			for _, v := range repos {
				fmt.Println(*v.Name, *v.SSHURL, *v.HTMLURL)
				allRepos = append(allRepos, repos...)

			}
		} else {
			break
		}
	}

	return allRepos, nil

}
