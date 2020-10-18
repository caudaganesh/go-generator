package runner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func GoImports(filepath string) error {
	cmd := exec.Command("goimports", "-w", filepath)
	buf := &bytes.Buffer{}
	cmd.Stderr = buf

	err := cmd.Run()

	stderrContent, _ := ioutil.ReadAll(buf)
	if len(stderrContent) != 0 {
		return fmt.Errorf("goimports failed: %s", string(stderrContent))
	}

	exitError, ok := err.(*exec.ExitError)
	if ok {
		return fmt.Errorf("goimports error: exit code: %d message: %s", exitError.ExitCode(), exitError.Error())
	}

	return err
}
