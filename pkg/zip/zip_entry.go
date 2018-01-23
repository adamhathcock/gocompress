package gocompress

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type zipFormatEntry struct {
	zipEntry *zip.File
}

func (entry zipFormatEntry) Name() string {
	if entry.zipEntry != nil {
		return entry.zipEntry.Name
	}
	return ""
}

func (entry zipFormatEntry) IsDirectory() bool {
	// just the suffix of '/' should be enough
	return (entry.zipEntry.CompressedSize64 == 0 && entry.zipEntry.UncompressedSize64 == 0 && strings.HasSuffix(entry.Name(), "/"))
}

func (entry zipFormatEntry) Mode() os.FileMode {
	if entry.zipEntry != nil {
		return entry.zipEntry.FileInfo().Mode()
	}
	return os.ModeAppend
}

func (entry *zipFormatEntry) Write(output io.Writer) error {
	if entry.zipEntry == nil {
		return errors.New("no Reader")
	}
	rc, err := entry.zipEntry.Open()
	if err != nil {
		return fmt.Errorf("%s: open compressed file: %v", entry.zipEntry.Name, err)
	}
	_, err = io.Copy(output, rc)
	return err
}
