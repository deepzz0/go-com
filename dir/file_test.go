package dir

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	src := "test/file/src/file_test.go"
	dst := "test/file/dst/file_test.go"
	if err := CopyFile(src, dst); err != nil {
		t.Errorf("CopyFile:%v", err)
	}
	if !IsExist(dst) {
		t.Log("file is not exist")
	}
	if size, err := FileSize(dst); err != nil {
		t.Errorf("FileSize:%v", err)
	} else {
		t.Log(fmt.Sprint(size))
	}
	if time, err := FileModTime(dst); err != nil {
		t.Errorf("FileModTime:%v", err)
	} else {
		t.Log(fmt.Sprint(time))
	}
	if size, err := FileSize(dst); err != nil {
		t.Errorf("FileSize:%v", err)
	} else {
		t.Log(HumaneFileSize(uint64(size)))
	}
}

func TestWriteFile(t *testing.T) {
	dst := "test/file/dst/file_test2.go"
	content := "package main \n import(\"fmt\")\nfunc main(){fmt.Println(\"hello word!！\")}"
	if err := WriteFile(dst, []byte(content)); err != nil {
		t.Errorf("WriteFile:%v", err)
	}
}
