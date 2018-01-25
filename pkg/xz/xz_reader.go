package xz

import (
	"fmt"
	"os"

	"github.com/adamhathcock/gocompress"
	"github.com/ulikunitz/xz"
)

// Reader is the entry point for using an archive reader on a Rar archive
var Reader xzFormatReader

type xzFormatReader struct {
	xzReader *xz.Reader
}

func (rfr *xzFormatReader) Close() error {
	return nil
}

func (rfr *xzFormatReader) OpenPath(path string) error {
	rf, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	r, err := xz.NewReader(rf)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	rfr.xzReader = r
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (rfr *xzFormatReader) ReadEntry() (gocompress.Entry, error) {
	return &xzFormatEntry{rfr.xzReader}, nil
}

func (rfr *xzFormatReader) ArchiveType() gocompress.ArchiveType {
	return gocompress.ZipArchive
}
