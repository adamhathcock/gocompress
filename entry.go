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
	CompressionType() CompressionType
}

type CompressionType int

const (
	None CompressionType = iota
	GZip
	BZip2
	PPMd
	Deflate
	Deflate64
	Rar
	LZMA
	BCJ
	BCJ2
	LZip
	Xz
	Unknown
)
