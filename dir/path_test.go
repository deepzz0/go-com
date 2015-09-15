package dir

import (
	"testing"
)

func TestPath(t *testing.T) {
	gopath := GetGOPATHs()
	t.Log(gopath)
	if appPath, err := GetSrcPath("github.com/smallchen0/com/dir"); err != nil {
		t.Errorf("GetSrcPath:%v", err)
	} else {
		t.Log(appPath)
	}

	if home, err := HomeDir(); err != nil {
		t.Errorf("HomeDir:%v", err)
	} else {
		t.Log(home)
	}
}
