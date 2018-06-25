package main

import (
	"os"
	"maridPoc/conf"
	"maridPoc/runbook"
	"maridPoc/util"
	"maridPoc/que"
)

var opsGenieCredentials conf.OpsGenieCredentials


func main() {
	opsGenieCredentials, _= conf.ReadCredentials()
	print("apiKey="+opsGenieCredentials.ApiKey+"\n")
	print("queueName="+opsGenieCredentials.QueueUrl +"\n")
	gitFlag := util.GetBoolEnv("GITFLAG")
	print(gitFlag)
	var config map[string]string
	if gitFlag {
		url := util.GetStrEnv("GITURL")
		username := util.GetStrEnv("GITUSERNAME")
		password := util.GetStrEnv("GITPASSWORD")
		if url == "" || username == "" || password == ""{
			panic("GITURL, GITUSERNAME or GITPASSWORD can not be empty")
		}
		config = conf.ReadConfigurationFromGit(url, username, password)
	} else {
		maridConfPath := os.Getenv("MARIDCONFPATH")
		print(maridConfPath)
		if maridConfPath == ""{
			maridConfPath = "/tmp/maridConf/.marid.conf"
		}
		config = conf.ReadLocalConfiguration(maridConfPath)
	}
	print("test")
	que.RunLoop(opsGenieCredentials, config)
	runbook.ExecuteRunbook("alertInfo", config)
}



