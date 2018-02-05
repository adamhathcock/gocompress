package rar

import (
	"errors"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
	"github.com/nwaples/rardecode"
)

type rarFormatEntry struct {
	rarReader *rardecode.Reader
	header    *rardecode.FileHeader
}

func (entry rarFormatEntry) Name() string {
	if entry.header != nil {
		return entry.header.Name
	}
	return ""
}

func (entry rarFormatEntry) IsDirectory() bool {
	if entry.header != nil {
		return entry.header.IsDir
	}
	return false
}

func (entry rarFormatEntry) Mode() os.FileMode {
	if entry.header != nil {
		return entry.header.Mode()
	}
	return os.ModeAppend
}

func (entry *rarFormatEntry) Write(output io.Writer) error {
	if entry.rarReader == nil {
		return errors.New("no Reader")
	}
	_, err := io.Copy(output, entry.rarReader)
	return err
}

func (entry rarFormatEntry) CompressionType() common.CompressionType {
	if entry.header != nil && entry.header.PackedSize == entry.header.UnPackedSize {
		return common.None
	}
	return common.Rar
}
