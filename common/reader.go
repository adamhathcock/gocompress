package common

import "io"

// Reader is a Generic Archive Reader interface
type Reader interface {
	Next() (Entry, error)
	ArchiveType() ArchiveType
}

type ReadCloser interface {
	Reader
	io.Closer
}
