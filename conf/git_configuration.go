package conf

import (
	"strings"
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	"io/ioutil"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"maridPoc/util"
)

func getDirectoryName(url string) (name string) {
	urlWithoutExtension := strings.TrimRight(url, ".git")
	lastIndex := strings.LastIndex(urlWithoutExtension, "/")
	if lastIndex == -1 {
		panic("Not a valid GITURL")
	}
	return urlWithoutExtension[lastIndex+1:]
}
func ReadConfigurationFromGit(url string, username string, password string) (map[string]string) {
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
	util.Check(err)
	os.Chdir(directoryName)
	files, err := ioutil.ReadDir(".")
	util.Check(err)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	conf := readConfigurations("/tmp/"+directoryName+"/.marid.conf")
	print("/tmp/"+directoryName+"\n")
	os.RemoveAll("/tmp/"+directoryName)
	return conf
}
