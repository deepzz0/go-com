package arc

import (
	"fmt"
	"testing"
)

func TestTarFileList(t *testing.T) {
	if str, err := TarFileList("test/arc.tar.bz2"); err != nil {
		t.Errorf("TarFileList:%v", err)
	} else {
		fmt.Println(str)
	}
	str, err := ArchiveFileList("test/arc.tar.gz")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)
}
