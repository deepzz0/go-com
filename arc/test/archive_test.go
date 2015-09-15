package arc

import (
	"testing"
)

func TestTarFileList(t *testing.T) {
	if str, err := TarFileList("test/archive.gz"); err != nil {
		t.Errorf("TarFileList:%v", err)
	} else {
		t.Log(str)
	}
}
