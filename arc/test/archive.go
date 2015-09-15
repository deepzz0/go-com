package arc

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ArchiveFileList(file string) ([]string, error) {
	if suffix := Suffix(file); suffix == ".gz" {
		return GzipFileList(file)
	} else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".tgz" {
		return TarFileList(file)
	} else if suffix == ".zip" {
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive.")
}

func GzipFileList(file string) ([]string, error) {

	return nil, nil
}

func TarFileList(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	r := tar.NewReader(f)
	files := make([]string, 0, 100)
	for {
		h, err := r.Next()
		if err != nil {
			break
		}
		files = append(files, h.Name)
	}

	return files, nil
}

func ZipFileList(file string) ([]string, error) {
	rc, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	files := make([]string, 0, len(rc.File))
	for _, f := range rc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		files = append(files, f.FileHeader.Name)
	}
	return files, nil
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
