package que

import (
	"maridPoc/conf"
	"maridPoc/runbook"
)


func RunLoop(credentials conf.OpsGenieCredentials, config map[string]string)  {
	apiKey := credentials.ApiKey
	queUrl := credentials.QueueUrl
	opsGenieToken, _ := checkTokenGenerateIfNeeded(nil, apiKey)
	var svc = createSQSClient(opsGenieToken)
	for {
		opsGenieToken, shouldGenerateClient := checkTokenGenerateIfNeeded(opsGenieToken, apiKey)
		if shouldGenerateClient {
			svc = createSQSClient(opsGenieToken)
		}
		if resp, err := readMessage(svc, &queUrl); err == nil {
			runbook.ExecuteRunbook(resp.String(), config)
			deleteMessage(svc, &queUrl)
		}
	}

}


