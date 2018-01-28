package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress"
)

// Reader is the entry point for using an archive reader on a Rar archive
var Reader zipFormatReader

type zipFormatReader struct {
	zipReader *zip.ReadCloser
	index     int
}

// IsZip checks the file has the Zip format signature by reading its beginning
// bytes and matching it against "PK\x03\x04"
func IsZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

func (zfr *zipFormatReader) Close() error {
	return zfr.zipReader.Close()
}

func (zfr *zipFormatReader) OpenPath(path string) error {
	var err error
	zfr.zipReader, err = zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (zfr *zipFormatReader) Next() (gocompress.Entry, error) {
	if zfr.index >= len(zfr.zipReader.File) {
		return nil, io.EOF
	}

	f := zfr.zipReader.File[zfr.index]

	zfr.index++
	return &zipFormatEntry{f}, nil
}

func (zfr *zipFormatReader) ArchiveType() gocompress.ArchiveType {
	return gocompress.ZipArchive
}
