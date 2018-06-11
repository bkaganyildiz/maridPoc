package main

import (
	"strconv"
	"fmt"
	"os"
	"maridPoc/conf"
)

var opsGenieCredentials conf.OpsGenieCredentials


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf(key))
	}
	return val
}

func getBoolEnv(key string) bool {
	val := getStrEnv(key)
	ret, err := strconv.ParseBool(val)
	check(err)
	return ret
}

func main() {
	opsGenieCredentials, _= conf.ReadCredentials()
	print("apiKey="+opsGenieCredentials.ApiKey+"\n")
	print("queueName="+opsGenieCredentials.QueueUrl +"\n")
	gitFlag := getBoolEnv("GITFLAG")
	print(gitFlag)
	if gitFlag {
		url := getStrEnv("GITURL")
		username := getStrEnv("GITUSERNAME")
		password := getStrEnv("GITPASSWORD")
		if url == "" || username == "" || password == ""{
			panic("GITURL, GITUSERNAME or GITPASSWORD can not be empty")
		}
		conf.ReadConfigurationFromGit(url, username, password)
	} else {
		maridConfPath := os.Getenv("MARIDCONFPATH")
		print(maridConfPath)
		if maridConfPath == ""{
			maridConfPath = "/tmp/maridConf/.marid.conf"
		}
		conf.ReadLocalConfiguration(maridConfPath)
	}
}



