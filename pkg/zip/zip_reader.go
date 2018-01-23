package gocompress

import (
	"archive/zip"
	"fmt"
	"github.com/adamhathcock/gocompress"
)

// ZipReader is the entry point for using an archive reader on a Rar archive
var ZipReader zipFormatReader

type zipFormatReader struct {
	zipReader *zip.ReadCloser
	index     int
}

func (rfr *zipFormatReader) Close() error {
	return nil
}

func (rfr *zipFormatReader) OpenPath(path string) error {
	var err error
	rfr.zipReader, err = zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (rfr *zipFormatReader) ReadEntry() (gocompress.Entry, error) {
	if rfr.index >= len(rfr.zipReader.File) {
		return gocompress.NilEntry, nil
	}

	f := rfr.zipReader.File[rfr.index]

	rfr.index++
	return &zipFormatEntry{f}, nil
}
