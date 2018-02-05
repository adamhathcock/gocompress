package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/adamhathcock/gocompress/common"
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

func (entry zipFormatEntry) CompressionType() common.CompressionType {
	switch ZipCompressionMethod(entry.zipEntry.Method) {
	case None:
		return common.None
	case Deflate:
		return common.Deflate
	case Deflate64:
		return common.Deflate64
	case BZip2:
		return common.BZip2
	case LZMA:
		return common.LZMA
	case PPMd:
		return common.PPMd
	default:
		return common.Unknown
	}
	return common.Rar
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
