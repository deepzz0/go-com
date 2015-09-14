package cmd

import (
	"github.com/smallchen0/go-com/log"
	"testing"
)

func TestExecCmd(t *testing.T) {
	str, str1, err := execCmd("", "open", "../log/log.go")
	log.Debugf(">>>%s", str)
	log.Debugf(">>>%s", str1)
	log.Debugf(">>>%#v", err)
}
