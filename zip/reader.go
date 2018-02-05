package zip

import (
	"archive/zip"
	"fmt"
	"io"

	"github.com/adamhathcock/gocompress/common"
)

func init() {
	zip.RegisterDecompressor(BZip2, BZip2Decompressor)
}

type Reader struct {
	zipReader *zip.ReadCloser
	index     int
}

func (reader *Reader) Close() error {
	return reader.zipReader.Close()
}

func OpenReader(path string) (common.ReadCloser, error) {
	var err error
	zfr := &Reader{}
	zfr.zipReader, err = zip.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("read: failed to create reader: %v", err)
	}
	return zfr, nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (reader *Reader) Next() (common.Entry, error) {
	if reader.index >= len(reader.zipReader.File) {
		return nil, io.EOF
	}

	f := reader.zipReader.File[reader.index]

	reader.index++
	return &zipFormatEntry{f}, nil
}

func (reader *Reader) ArchiveType() common.ArchiveType {
	return common.ZipArchive
}
