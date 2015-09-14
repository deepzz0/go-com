package main

import (
	"github.com/smallchen0/go-com/arc"
	"github.com/smallchen0/go-com/log"
)

func main() {
	if str, err := arc.TarFileList("arc/test/2.gz"); err != nil {
		log.Errorf("TarFileList:%v", err)
	} else {
		log.Debug(str)
	}
}
