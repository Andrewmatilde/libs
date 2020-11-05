package crawler

import (
	"context"
	"github.com/google/go-github/v32/github"
	"log"
)

func ListTags(client *github.Client, owner string, repo string) []string {
	listOpt := github.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	LastPage := 1
	var tagNames []string
	for LastPage != 0 {
		tags, rsp, err := client.Repositories.ListTags(context.Background(), owner, repo, &listOpt)
		if err != nil {
			log.Println(err)
			return nil
		}
		listOpt.Page = rsp.NextPage
		LastPage = rsp.LastPage
		for _, tag := range tags {
			tagNames = append(tagNames, *tag.Name)
		}
	}
	return tagNames
}
