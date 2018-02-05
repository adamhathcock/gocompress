package gocompress

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
	"github.com/adamhathcock/gocompress/rar"
	"github.com/adamhathcock/gocompress/tar"
	"github.com/adamhathcock/gocompress/zip"
)

// OpenFilePath opens a specific path to a supported archive and returns a Reader
func OpenReader(path string) (common.Reader, error) {
	if rar.IsRar(path) {
		reader, err := rar.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	if zip.IsZip(path) {
		reader, err := zip.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	if tar.IsTar(path) {
		reader, err := tar.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	return nil, errors.New(path + " has no valid format detected")
}

// Extract will extract all files from the source archive path to a destination folder
func Extract(source string, destination string) error {
	reader, err := OpenReader(source)
	if err != nil {
		return err
	}
	var entry common.Entry
	for {
		entry, err = reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		path := filepath.Join(destination, entry.Name())
		internal.WriteNewFile(path, 666)
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
