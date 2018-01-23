package gocompress

import (
	"io"
)

// ArchiveReader is a Generic Archive Reader interface
type ArchiveReader interface {
	OpenPath(path string) error
	Open(io.Reader) error
	ReadEntry() (Entry, error)
	Close() error
}
