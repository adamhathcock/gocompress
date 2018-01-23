package gocompress

import (
	"io"
	"os"
)

// Entry is the generic archive entry interface when reading archives
type Entry interface {
	Name() string
	IsDirectory() bool
	Mode() os.FileMode
	Write(output io.Writer) error
}
