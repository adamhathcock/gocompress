package common

import "io"

// Reader is a Generic Archive Reader interface
type Reader interface {
	io.Closer
	Next() (Entry, error)
	ArchiveType() ArchiveType
}

// ArchiveType enum
type ArchiveType int

const (
	// Rar ArchiveType
	RarArchive ArchiveType = iota
	// Zip ArchiveType
	ZipArchive
	// Tar ArchiveType
	TarArchive
)
