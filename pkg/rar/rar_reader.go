package rar

import (
	"fmt"
	"io"
	"os"

	"github.com/nwaples/rardecode"
	"github.com/adamhathcock/gocompress"
)

// RarReader is the entry point for using an archive reader on a Rar archive
var RarReader rarFormatReader

type rarFormatReader struct {
	rarReader *rardecode.Reader
}

func (rfr *rarFormatReader) Close() error {
	return nil
}

func (rfr *rarFormatReader) OpenPath(path string) error {
	rf, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}

	return rfr.Open(rf)
}

func (rfr *rarFormatReader) Open(input io.Reader) error {
	var err error
	rfr.rarReader, err = rardecode.NewReader(input, "")
	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (rfr *rarFormatReader) ReadEntry() (gocompress.Entry, error) {
	header, err := rfr.rarReader.Next()
	if err == io.EOF {
		return gocompress.NilEntry, nil
	} else if err != nil {
		return gocompress.NilEntry, err
	}

	return &rarFormatEntry{
		rfr.rarReader,
		header}, nil
}
