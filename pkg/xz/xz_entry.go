package xz

import (
	"io"
	"os"

	"github.com/ulikunitz/xz"
)

type xzFormatEntry struct {
	xzReader *xz.Reader
}

func (entry xzFormatEntry) Name() string {
	return ""
}

func (entry xzFormatEntry) IsDirectory() bool {
	// just the suffix of '/' should be enough
	return false
}

func (entry xzFormatEntry) Mode() os.FileMode {
	return os.ModeAppend
}

func (entry *xzFormatEntry) Write(output io.Writer) error {
	_, err := io.Copy(output, entry.xzReader)
	return err
}
