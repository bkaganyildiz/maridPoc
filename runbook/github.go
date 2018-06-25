package runbook

import (
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"io/ioutil"
	"context"
)



func getRunbookFromGithub(owner string, repo string, filepath string, token string) (string, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	runbook, _ := client.Repositories.DownloadContents(context.Background(), owner, repo, filepath, nil)
	defer runbook.Close()
	bytes, _ := ioutil.ReadAll(runbook)
	return string(bytes), nil
}
