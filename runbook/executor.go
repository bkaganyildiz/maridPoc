package runbook

import (
	"maridPoc/util"
)

var GIT_OWNER = "testrunbook"
var GIT_REPO = "runbooks"
var GIT_PATH = "test.sh"

func ExecuteRunbook(alertInfo string, config map[string]string) {
	runbook, _ := getRunbookFromGithub(GIT_OWNER, GIT_REPO, GIT_PATH, util.GetStrEnv("RUNBOOKTOKEN"))
	filename, _ := generateTempFile(&runbook)
	commandExecutor(filename, alertInfo, util.MapToQueryParamConverter(config))
}

