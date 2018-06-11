package conf

import (
	"os/user"
	"fmt"
	"io/ioutil"
	"strings"
	"errors"
	"net/http"
	"bytes"
	"encoding/json"
)
const (
	createQueue = "http://localhost:5000/queue-generator"
	dummyMessageCreator = "http://localhost:5000/message-publisher"
)

type OpsGenieCredentials struct {
	ApiKey   string
	QueueUrl string
}

func ReadCredentials() (opsGenieCredentials OpsGenieCredentials, err error) {
	apiKey, err := readApiKey()
	queUrl := createQueueAndGetItsName(apiKey)
	createDummyMessages(queUrl)
	credentials := OpsGenieCredentials{
		apiKey,
		queUrl,
	}
	return credentials, err
}
func createDummyMessages(queUrl string) {
	body := map[string]string{"queueUrl": queUrl}
	jsonValue, _ := json.Marshal(body)
	resp,  _ := http.Post(dummyMessageCreator, "application/json", bytes.NewBuffer(jsonValue))
	if resp.StatusCode == 200 {
		print("\nyey\n")
	}
}
func createQueueAndGetItsName(apiKey string) (queName string) {
	body := map[string]string{"apiKey": apiKey}
	jsonValue, _ := json.Marshal(body)
	resp,  _ := http.Post(createQueue, "application/json", bytes.NewBuffer(jsonValue))
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		payload, _ := ioutil.ReadAll(resp.Body)
		m := make(map[string]string)
		err := json.Unmarshal(payload, &m)
		check(err)
		return m["queue_name"]
	} else {
		defer resp.Body.Close()
		payload, _ := ioutil.ReadAll(resp.Body)
		m := make(map[string]string)
		err := json.Unmarshal(payload, &m)
		check(err)
		panic(m["_error"])
	}
	return ""
}
func readApiKey() (apiKey string, err error){
	usr, err := user.Current()
	fmt.Println( usr.HomeDir )
	dat, err := ioutil.ReadFile(usr.HomeDir+"/.opsgenie/credentials")
	out, err := after(string(dat))
	return out, err
}

func after(value string) (string, error){
	// Get substring after a string.
	a := "apiKey="
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return "", errors.New("ApiKey must be provided");
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return "", errors.New("ApiKey must be provided");
	}
	return value[adjustedPos:len(value)-1], nil
}