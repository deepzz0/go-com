package cmd

import (
	"github.com/deepzz0/go-com/log"
	"testing"
)

func TestExecCmd(t *testing.T) {
	// str, str1, err := execCmd("", "open", "../log/log.go")
	// log.Debugf(">>>%s", str)
	// log.Debugf(">>>%s", str1)
	// log.Debugf(">>>%#v", err)

	// mysqldump -uroot -proot gamecenter bag --where="" > /Users/chen/a.sql
	str, str1, _ := execCmd("", "mysqldump", "-uroot", "-proot", "gamecenter", "test", "--where=id>1 and id<5")
	log.Debugf(">>>%s,>>>>%s,%s", str, str1)
}
