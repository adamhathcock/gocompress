package tar

import (
	"archive/tar"
	"errors"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
)

type tarFormatEntry struct {
	tarReader   *tar.Reader
	header      *tar.Header
	compression common.CompressionType
}

func (entry tarFormatEntry) Name() string {
	if entry.header != nil {
		return entry.header.Name
	}
	return ""
}

func (entry tarFormatEntry) IsDirectory() bool {
	if entry.header != nil {
		return entry.header.Typeflag == tar.TypeDir
	}
	return false
}

func (entry tarFormatEntry) Mode() os.FileMode {
	if entry.header != nil {
		return entry.header.FileInfo().Mode()
	}
	return os.ModeAppend
}

func (entry *tarFormatEntry) Write(output io.Writer) error {
	if entry.tarReader == nil {
		return errors.New("no Reader")
	}
	_, err := io.Copy(output, entry.tarReader)
	return err
}

func (entry tarFormatEntry) CompressionType() common.CompressionType {
	return entry.compression
}
