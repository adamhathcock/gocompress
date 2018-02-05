package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
)

func init() {
	zip.RegisterDecompressor(BZip2, BZip2Decompressor)
}

type Reader struct {
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

func (zfr *Reader) Close() error {
	return zfr.zipReader.Close()
}

func OpenReader(path string) (common.Reader, error) {
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
func (zfr *Reader) Next() (common.Entry, error) {
	if zfr.index >= len(zfr.zipReader.File) {
		return nil, io.EOF
	}

	f := zfr.zipReader.File[zfr.index]

	zfr.index++
	return &zipFormatEntry{f}, nil
}

func (zfr *Reader) ArchiveType() common.ArchiveType {
	return common.ZipArchive
}
