package common

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
