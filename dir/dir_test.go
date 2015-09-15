package dir

import (
	// "fmt"
	// "strings"
	"testing"
)

func TestDir(t *testing.T) {
	path := "test"
	if !IsDir(path) {
		t.Errorf("IsDir:%s", path)
	}
}

var f = func(filepath string) bool {
	// if str := fmt.Sprintf("test/srcdir/%s", filepath); strings.Contains(str, ".go") {
	// 	return true
	// }

	return false
}

func TestCopyDir(t *testing.T) {
	srcpath := "test/dir/srcdir"
	dstpath := "test/dir/dstdir"

	if err := CopyDir(srcpath, dstpath, f); err != nil {
		t.Errorf("CopyDir: %v", err)
	}
}

func TestStatDir(t *testing.T) {
	rootpath := "test/dir/srcdir"

	if str, err := StatDir(rootpath, false); err != nil {
		t.Errorf("StatDir:%v", err)
	} else {
		t.Log(str)
	}
}

func TestGetAllSubDirs(t *testing.T) {
	rootpath := "test/dir/srcdir"
	if str, err := GetAllSubDirs(rootpath); err != nil {
		t.Errorf("GetAllSubDirs:%v", err)
	} else {
		t.Log(str)
	}
}

func TestGetFileListBySuffix(t *testing.T) {
	rootpath := "test/dir/srcdir"
	if str, err := GetFileListBySuffix(rootpath, ".go"); err != nil {
		t.Errorf("GetAllSubDirs:%v", err)
	} else {
		t.Log(str)
	}
}
