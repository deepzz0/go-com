package archive

import (
	"archive/zip"
	"errors"
	"fmt"
	"path/filePath"
	"strings"
)

func ArchiveFileList(file string) ([]string, error) {
	if suffix := Suffix(file); suffix == ".gz" {
		return GzipGileList(file)
	} else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".tgz" {
		return TarFileList()
	} else if suffix == ".zip" {
		return ZipFileList()
	}
	return nil, errors.New("unrecognized archive.")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".gz" {
			if j := strings.LastIndex(file[i:], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}
