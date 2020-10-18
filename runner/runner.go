package runner

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

func Run(filename string) (io.Reader, error) {
	stdoutBuffer := &bytes.Buffer{}
	stderrBuffer := &bytes.Buffer{}

	cmd := exec.Command("go", "run", filename)
	cmd.Stdout = stdoutBuffer
	cmd.Stderr = stderrBuffer

	cmdErr := cmd.Run()

	// check stderr first if there are stderr return error
	// ignore the error from read all (possibly caused by
	// memory allocation
	stderrContent, _ := ioutil.ReadAll(stderrBuffer)
	if len(stderrContent) != 0 {
		return nil, fmt.Errorf("Run %s error : %s", GetFullPath(), string(stderrContent))
	}

	// if cmdErr then return stdout to error
	if cmdErr != nil {
		out, _ := ioutil.ReadAll(stdoutBuffer)
		return nil, fmt.Errorf("Run %s error stdout: %s", GetFullPath(), string(out))
	}

	// else return stdout
	return stdoutBuffer, nil
}
