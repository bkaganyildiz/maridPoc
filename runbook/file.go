package runbook

import (
	"path/filepath"
	"strings"
	"os"
)

func generateTempFile(runbook *string) (string, error) {
	fileName, _ := filepath.Abs(filepath.Join("/tmp", strings.Join([]string{"test", ".sh"}, "")))
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
		return fileName, err
	} else {
		// Make the file executable.
		os.Chmod(fileName, os.ModePerm)
		defer f.Close()
	}
	_, err = f.WriteString(*runbook)
	return fileName, nil
}
