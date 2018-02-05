package rar

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
	"github.com/nwaples/rardecode"
)

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

type Reader struct {
	rarReader *rardecode.Reader
}

func (rfr *Reader) Close() error {
	return nil
}

func OpenReader(path string) (common.Reader, error) {
	rf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
	}

	return Open(rf)
}

func Open(input io.Reader) (common.Reader, error) {
	var err error
	rfr := &Reader{}
	rfr.rarReader, err = rardecode.NewReader(input, "")
	if err != nil {
		return nil, fmt.Errorf("read: failed to create reader: %v", err)
	}
	return rfr, nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (rfr *Reader) Next() (common.Entry, error) {
	header, err := rfr.rarReader.Next()
	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return &rarFormatEntry{
		rfr.rarReader,
		header}, nil
}

func (rfr *Reader) ArchiveType() common.ArchiveType {
	return common.RarArchive
}
