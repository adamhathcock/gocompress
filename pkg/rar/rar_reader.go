package rar

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress"
	"github.com/nwaples/rardecode"
)

// Reader is the entry point for using an archive reader on a Rar archive
var Reader rarFormatReader

// IsRar checks the file has the RAR 1.5 or 5.0 format signature by reading its
// beginning bytes and matching it
func IsRar(rarPath string) bool {
	f, err := os.Open(rarPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 8)
	if n, err := f.Read(buf); err != nil || n < 8 {
		return false
	}

	return bytes.Equal(buf[:7], []byte("Rar!\x1a\x07\x00")) || // ver 1.5
		bytes.Equal(buf, []byte("Rar!\x1a\x07\x01\x00")) // ver 5.0
}

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
