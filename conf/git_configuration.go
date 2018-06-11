package conf

import (
	"strings"
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	"io/ioutil"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDirectoryName(url string) (name string) {
	urlWithoutExtension := strings.TrimRight(url, ".git")
	lastIndex := strings.LastIndex(urlWithoutExtension, "/")
	if lastIndex == -1 {
		panic("Not a valid GITURL")
	}
	return urlWithoutExtension[lastIndex+1:]
}
func ReadConfigurationFromGit(url string, username string, password string)  {
	os.Chdir("/tmp")
	directoryName := getDirectoryName(url)
	_, err := git.PlainClone(directoryName, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth: &http.BasicAuth{
			username,
			password,
		},
	})
	check(err)
	os.Chdir(directoryName)
	files, err := ioutil.ReadDir(".")
	check(err)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	readConfigurations("/tmp/"+directoryName+"/.marid.conf")
	print("/tmp/"+directoryName)
	os.RemoveAll("/tmp/"+directoryName)
}
