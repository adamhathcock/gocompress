package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/adamhathcock/gocompress"
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
	return entry.zipEntry.CompressedSize64 == 0 && entry.zipEntry.UncompressedSize64 == 0 && strings.HasSuffix(entry.Name(), "/")
}

func (entry zipFormatEntry) Mode() os.FileMode {
	if entry.zipEntry != nil {
		return entry.zipEntry.FileInfo().Mode()
	}
	return os.ModeAppend
}

func (entry zipFormatEntry) Write(output io.Writer) error {
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

func (entry zipFormatEntry) CompressionType() gocompress.CompressionType {
	switch ZipCompressionMethod(entry.zipEntry.Method) {
	case None:
		return gocompress.None
	case Deflate:
		return gocompress.Deflate
	case Deflate64:
		return gocompress.Deflate64
	case BZip2:
		return gocompress.BZip2
	case LZMA:
		return gocompress.LZMA
	case PPMd:
		return gocompress.PPMd
	default:
		return gocompress.Unknown
	}
	return gocompress.Rar
}

type ZipCompressionMethod uint16

const (
	None      ZipCompressionMethod = 0
	Deflate                        = 8
	Deflate64                      = 9
	BZip2                          = 12
	LZMA                           = 14
	PPMd                           = 98
	WinzipAes                      = 0x63 //http://www.winzip.com/aes_info.htm
)
