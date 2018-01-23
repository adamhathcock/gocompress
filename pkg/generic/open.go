package generic

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/adamhathcock/gocompress"
	"github.com/adamhathcock/gocompress/pkg/rar"
	"github.com/adamhathcock/gocompress/pkg/zip"
)

// OpenFilePath opens a specific path to a supported archive and returns a Reader
func OpenFilePath(path string) (gocompress.Reader, error) {
	if rar.IsRar(path) {
		rr := rar.Reader
		err := rr.OpenPath(path)
		if err != nil {
			return nil, err
		}
		// why do I have to do this?
		s := &rr
		return s, nil
	}
	if zip.IsZip(path) {
		zr := zip.Reader
		err := zr.OpenPath(path)
		if err != nil {
			return nil, err
		}
		return &zr, nil
	}
	return nil, errors.New(path + " has no valid format detected")
}

// Extract will extract all files from the source archive path to a destination folder
func Extract(source string, destination string) error {
	reader, err := OpenFilePath(source)
	if err != nil {
		return err
	}
	var entry gocompress.Entry
	for {
		entry, err = reader.ReadEntry()
		if err != nil {
			return err
		}
		if entry == nil {
			break
		}
		path := filepath.Join(destination, entry.Name())
		writeNewFile(path, 666)
		writer, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
		err = entry.Write(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeNewFile(fpath string, fm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fpath), 0755)
	if err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}

	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close()

	err = out.Chmod(fm)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}

	return nil
}
