package runbook

import (
	"os/exec"
	"bytes"
)

func commandExecutor(fileName string, alertInfo string, config string)  {
	print("\n" + fileName)
	cmd := exec.Command(fileName, alertInfo, config)
	cmdOutput := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdErr
	err := cmd.Run()
	output := cmdOutput.String()
	print(output)
	print(cmdErr.String())
	if err != nil {
		print(err.Error())
	}
}
