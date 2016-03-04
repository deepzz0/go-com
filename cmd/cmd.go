package cmd

import (
	"bytes"
	"os/exec"
)

func execCmd(dir string, cmdName string, args ...string) ([]byte, []byte, error) {
	cmd := exec.Command(cmdName, args...)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Dir = dir
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func ExecCmd(cmdName string, args ...string) ([]byte, []byte, error) {
	return execCmd("", cmdName, args...)
}

func ExecCmdDir(dir string, cmdName string, args ...string) (string, string, error) {
	stdout, stderr, err := execCmd(dir, cmdName, args...)
	return string(stdout), string(stderr), err
}
