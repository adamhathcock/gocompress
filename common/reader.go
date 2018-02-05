package common

import "io"

// Reader is a Generic Archive Reader interface
type Reader interface {
	io.Closer
	Next() (Entry, error)
	ArchiveType() ArchiveType
}
