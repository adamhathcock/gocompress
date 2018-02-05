package common

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
