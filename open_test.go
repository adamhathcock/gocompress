package gocompress

import (
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
)

func TestOpenReader_Rar(t *testing.T) {
	internal.ReadCloserExtractionTest(t, "files/archives/rar/Rar.rar", OpenReader, common.RarArchive, common.Rar)
}

func TestOpenReader_Zip(t *testing.T) {
	internal.ReadCloserExtractionTest(t, "files/archives/zip/Zip.zip", OpenReader, common.ZipArchive, common.Deflate)
}

func TestOpenReader_Tar(t *testing.T) {
	internal.ReadCloserExtractionTest(t, "files/archives/tar/Tar.tar", OpenReader, common.TarArchive, common.None)
}

/*

func TestOpenReader_Rar(t *testing.T) {
	internal.ReaderExtractionTest(t, "files/archives/rar/Rar.rar", NewReader, common.RarArchive, common.Rar)
}

func TestOpenReader_Zip(t *testing.T) {
	internal.ReaderExtractionTest(t, "files/archives/zip/Zip.zip", NewReader, common.ZipArchive, common.Deflate)
}

func TestOpenReader_Tar(t *testing.T) {
	internal.ReaderExtractionTest(t, "files/archives/tar/Tar.tar", NewReader, common.TarArchive, common.None)
}*/
