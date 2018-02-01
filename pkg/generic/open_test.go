package generic

import (
	"path/filepath"
	"testing"

	"github.com/adamhathcock/gocompress"
	"github.com/stretchr/testify/require"
)

func TestOpenFilePath_Rar(t *testing.T) {
	require := require.New(t)

	aPath, _ := filepath.Abs("../../files/archives/rar/Rar.rar")
	reader, err := OpenFilePath(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)
	gocompress.GenericExtractionTest(require, reader, gocompress.RarArchive, gocompress.Rar)
}

func TestOpenFilePath_Zip(t *testing.T) {
	require := require.New(t)

	aPath, _ := filepath.Abs("../../files/archives/zip/Zip.zip")
	reader, err := OpenFilePath(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)
	gocompress.GenericExtractionTest(require, reader, gocompress.ZipArchive, gocompress.Deflate)
}

func TestOpenFilePath_Tar(t *testing.T) {
	require := require.New(t)

	aPath, _ := filepath.Abs("../../files/archives/tar/Tar.tar")
	reader, err := OpenFilePath(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)
	gocompress.GenericExtractionTest(require, reader, gocompress.TarArchive, gocompress.None)
}
