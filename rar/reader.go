package rar

import (
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
	"github.com/nwaples/rardecode"
)

type Reader struct {
	rarReader *rardecode.Reader
}

type ReadCloser struct {
	Reader
	closer io.ReadCloser
}

func (rfr *ReadCloser) Close() error {
	return rfr.closer.Close()
}

func OpenReader(path string) (common.ReadCloser, error) {
	rf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	reader, err := open(rf)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	rc := new(ReadCloser)
	rc.rarReader = reader
	rc.closer = rf
	return rc, nil
}

func Open(input io.Reader) (common.Reader, error) {
	reader, err := open(input)
	if err != nil {
		return nil, err
	}
	return &Reader{reader}, nil
}

func open(input io.Reader) (*rardecode.Reader, error) {
	rarReader, err := rardecode.NewReader(input, "")
	if err != nil {
		return nil, fmt.Errorf("read: failed to create reader: %v", err)
	}
	return rarReader, nil
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

	return &rarEntry{
		rfr.rarReader,
		header}, nil
}

func (rfr *Reader) ArchiveType() common.ArchiveType {
	return common.RarArchive
}
