package xz

import (
	"fmt"
	"os"

	"github.com/adamhathcock/gocompress"
	"github.com/ulikunitz/xz"
)

type Reader struct {
	xzReader *xz.Reader
}

func (rfr *Reader) Close() error {
	return nil
}

func (rfr *Reader) OpenPath(path string) error {
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

func (rfr *Reader) Next() (gocompress.Entry, error) {
	return &xzFormatEntry{rfr.xzReader}, nil
}

func (rfr *Reader) ArchiveType() gocompress.ArchiveType {
	return gocompress.ZipArchive
}
