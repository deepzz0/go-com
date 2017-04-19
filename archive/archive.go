package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// list file by archive
var FunctionForSuffix = map[string]func(string) ([]string, error){
	".gz":      GzipFileList,
	".tar":     TarFileList,
	".tar.gz":  TarFileList,
	".tar.bz2": TarFileList,
	".tgz":     TarFileList,
	".zip":     ZipFileList,
}

// archive file list
func ArchiveFileList(file string) ([]string, error) {
	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
		return function(file)
	}

	return nil, errors.New("unrecognized archive")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		// .bz2, .gz, .xz
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			// .tar.gz
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}

		return file[i:]
	}

	return file
}

// zip file list
func ZipFileList(filename string) ([]string, error) {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()

	var files []string
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() || strings.Contains(file.Name, ".DS_Store") {
			continue
		}
		files = append(files, file.Name)
	}

	return files, nil
}

// gzip file list
func GzipFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	return []string{gzipReader.Header.Name}, nil
}

// tar file list
func TarFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var tarReader *tar.Reader
	if strings.HasSuffix(filename, ".gz") ||
		strings.HasSuffix(filename, ".tgz") {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		defer gzipReader.Close()

		tarReader = tar.NewReader(gzipReader)
	} else if strings.HasSuffix(filename, ".bz2") {
		bz2Reader := bzip2.NewReader(reader)
		tarReader = tar.NewReader(bz2Reader)
	} else {
		tarReader = tar.NewReader(reader)
	}

	var files []string
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return files, err
		}

		if header == nil {
			break
		}

		if header.FileInfo().IsDir() || strings.Contains(header.Name, ".DS_Store") {
			continue
		}
		files = append(files, header.Name)
	}

	return files, nil
}

// unpack archive
func UnpackArchive(filename string) error {
	return UnpackArchive2Path(filename, "")
}

// unpack archive to path
func UnpackArchive2Path(filename string, path string) error {
	if suffix := Suffix(filename); suffix == ".zip" {
		return UnpackZip(filename, path)
	} else if suffix == ".tar.bz2" ||
		suffix == ".tar.gz" ||
		suffix == ".tar" {
		return UnpackTar(filename, path)
	}

	return errors.New("unrecognize archive.")
}

// unpack zip
func UnpackZip(filename string, path string) (err error) {
	var reader *zip.ReadCloser
	if reader, err = zip.OpenReader(filename); err != nil {
		return err
	}
	defer reader.Close()

	for _, zipFile := range reader.Reader.File {
		filename := sanitizedName(zipFile.Name, path)
		if strings.HasSuffix(zipFile.Name, "/") ||
			strings.HasSuffix(zipFile.Name, "\\") {
			if err = os.MkdirAll(filename, 0755); err != nil {
				return err
			}
		} else {
			if err = unpackZippedFile(filename, zipFile); err != nil {
				return err
			}
		}
	}

	return nil
}

func unpackZippedFile(filename string, zipFile *zip.File) (err error) {
	var writer *os.File
	if writer, err = os.Create(filename); err != nil {
		return err
	}
	defer writer.Close()

	var reader io.ReadCloser
	if reader, err = zipFile.Open(); err != nil {
		return err
	}
	defer reader.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	if filename == zipFile.Name {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, zipFile.Name)
	}

	return nil
}

// unpack tar
func UnpackTar(filename string, path string) (err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return err
	}
	defer file.Close()

	var fileReader io.Reader = file
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		if decompressor, err = gzip.NewReader(file); err != nil {
			return err
		}
		defer decompressor.Close()
	} else if strings.HasSuffix(filename, ".bz2") {
		fileReader = bzip2.NewReader(file)
	}

	var reader *tar.Reader
	if decompressor != nil {
		reader = tar.NewReader(decompressor)
	} else {
		reader = tar.NewReader(fileReader)
	}

	return unpackTarFiles(reader, path)
}

func unpackTarFiles(reader *tar.Reader, path string) (err error) {
	var header *tar.Header
	for {
		if header, err = reader.Next(); err != nil {
			if err == io.EOF {
				return nil // OK
			}
			return err
		}

		filename := sanitizedName(header.Name, path)
		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.MkdirAll(filename, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err = unpackTarFile(filename, header.Name, reader); err != nil {
				return err
			}
		}
	}

	return nil
}

func unpackTarFile(filename, tarFilename string, reader *tar.Reader) (err error) {
	var writer *os.File
	if writer, err = os.Create(filename); err != nil {
		return err
	}
	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	if filename == tarFilename {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, tarFilename)
	}

	return nil
}

func sanitizedName(filename string, path string) string {
	if len(filename) > 1 && filename[1] == ':' {
		filename = filename[2:]
	}
	filename = strings.TrimLeft(filename, "\\/.")
	filename = strings.Replace(filename, "../", "", -1)
	if path == "" {
		return strings.Replace(filename, "..\\", "", -1)
	}
	return fmt.Sprintf("%s/%s", path, filename)
}

// create archive.
func CreateZip(filename string, files []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	zipper := zip.NewWriter(file)
	for _, name := range files {
		if err := writeFile2Zip(zipper, name); err != nil {
			return err
		}
	}

	return nil
}

func writeFile2Zip(zipper *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = sanitizedName(filename, "")

	writer, err := zipper.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)

	return err
}

func CreateTar(filename string, files []string) error {
	if !strings.Contains(filename, ".tar") {
		filename = filename + ".tar"
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileWriter io.WriteCloser = file
	if strings.HasSuffix(filename, ".gz") {
		defer fileWriter.Close()
	}
	writer := tar.NewWriter(fileWriter)
	defer writer.Close()

	for _, name := range files {
		if err := writeFile2Tar(writer, name); err != nil {
			return err
		}
	}

	return nil
}

func writeFile2Tar(writer *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    sanitizedName(filename, ""),
		Mode:    int64(stat.Mode()),
		Uid:     os.Getuid(),
		Gid:     os.Getgid(),
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
	}
	if err = writer.WriteHeader(header); err != nil {
		return err
	}
	_, err = io.Copy(writer, file)

	return err
}
