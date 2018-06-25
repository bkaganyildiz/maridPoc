package conf

import (
	"io/ioutil"
	"strings"
)

func readConfigurations(path string) (conf map[string]string) {
	print(path+"\n")
	configuration := make(map[string]string)
	dat, _ := ioutil.ReadFile(path)
	confs := strings.Split(string(dat), "\n")
	for _, conf := range confs {
		equalIndex := strings.Index(conf, "=")
		if equalIndex == -1 {
			continue
		}
		configuration[conf[:equalIndex]] = conf[equalIndex+1:]
	}
	return configuration
}
