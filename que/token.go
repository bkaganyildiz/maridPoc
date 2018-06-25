package que

import (
	"encoding/json"
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
	"errors"
	"maridPoc/util"
)

const (
	TOKEN_ENDPOINT = "http://localhost:5000/sts-generator"
 	REGION  = "eu-west-1"
)

type OpsGenieToken struct {
	akid string
	token string
	secret string
	timestamp time.Time
}

func checkTokenGenerateIfNeeded(opsGenieToken *OpsGenieToken, apiKey string) (*OpsGenieToken, bool) {
	isTime := checkTimeForTokenGenerations(opsGenieToken)
	if  !isTime || opsGenieToken == nil {
		opsGenieToken, err := callForAnotherToken(apiKey)
		util.Check(err)
		return opsGenieToken, isTime
	} else {
		return opsGenieToken, isTime
	}
}

func callForAnotherToken(apiKey string) (*OpsGenieToken, error) {
	body := map[string]string{"apiKey": apiKey}
	jsonValue, _ := json.Marshal(body)
	resp,  _ := http.Post(TOKEN_ENDPOINT, "application/json", bytes.NewBuffer(jsonValue))
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		payload, _ := ioutil.ReadAll(resp.Body)
		m := make(map[string]string)
		_ = json.Unmarshal(payload, &m)
		return &OpsGenieToken{
			m["AccessKeyId"],
			m["SessionToken"],
			m["SecretAccessKey"],
			time.Now(),
		}, nil
	} else {
		defer resp.Body.Close()
		payload, _ := ioutil.ReadAll(resp.Body)
		m := make(map[string]string)
		_ = json.Unmarshal(payload, &m)
		return nil, errors.New(m["_error"])
	}
}

func checkTimeForTokenGenerations(opsgenieToken *OpsGenieToken) (bool) {
	if opsgenieToken == nil {
		return true
	}
	difference := time.Since(opsgenieToken.timestamp)
	if difference.Minutes() >= 1 {
		return true
	}
	return false
}
