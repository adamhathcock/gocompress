package gocompress

// Reader is a Generic Archive Reader interface
type Reader interface {
	OpenPath(path string) error
	ReadEntry() (Entry, error)
	Close() error
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
