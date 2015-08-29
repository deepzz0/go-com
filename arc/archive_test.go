package arc

import (
	"testing"
)

func TestTarFileList(t *testing.T) {
	if str, err := TarFileList("test/go1.5beta1.src.tar.gz"); err != nil {
		t.Errorf("TarFileList:%v", err)
	} else {
		t.Log(str)
	}
}
