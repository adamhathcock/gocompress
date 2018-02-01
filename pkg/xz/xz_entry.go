package xz

import (
	"io"
	"os"

	"github.com/adamhathcock/gocompress"
	"github.com/ulikunitz/xz"
)

type xzFormatEntry struct {
	xzReader *xz.Reader
}

func (entry xzFormatEntry) Name() string {
	return ""
}

func (entry xzFormatEntry) IsDirectory() bool {
	return false
}

func (entry xzFormatEntry) Mode() os.FileMode {
	return os.ModeAppend
}

func (entry *xzFormatEntry) Write(output io.Writer) error {
	_, err := io.Copy(output, entry.xzReader)
	return err
}
func (entry *xzFormatEntry) CompressionType() gocompress.CompressionType {
	return gocompress.Xz
}
