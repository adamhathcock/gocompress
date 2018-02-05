package gocompress

import (
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
)

func TestOpenFilePath_Rar(t *testing.T) {
	internal.ExtractionTest(t, "files/archives/rar/Rar.rar", OpenReader, common.RarArchive, common.Rar)
}

func TestOpenFilePath_Zip(t *testing.T) {
	internal.ExtractionTest(t, "files/archives/zip/Zip.zip", OpenReader, common.ZipArchive, common.Deflate)
}

func TestOpenFilePath_Tar(t *testing.T) {
	internal.ExtractionTest(t, "files/archives/tar/Tar.tar", OpenReader, common.TarArchive, common.None)
}
