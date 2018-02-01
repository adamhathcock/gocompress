package gocompress

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func GenericExtractionTest(require *require.Assertions, reader Reader, archiveType ArchiveType, compressionType CompressionType) {
	extract(require, reader, archiveType, compressionType)
}

func ExtractionTest(t *testing.T, reader Reader, archive string, archiveType ArchiveType, compressionType CompressionType) {
	require := require.New(t)

	aPath, err := filepath.Abs("../../files/archives/" + archive)
	err = reader.OpenPath(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)
	extract(require, reader, archiveType, compressionType)
}

func extract(require *require.Assertions, reader Reader, archiveType ArchiveType, compressionType CompressionType) {

	tmp, err := MakeTempDir(".")

	require.Nil(err, "Could not create temp dir %v", err)

	tmp, err = filepath.Abs(tmp)
	defer os.RemoveAll(tmp)

	require.Nil(err, "Could not read abs extracted path\n\t %v", err)

	extracted, err := filepath.Abs("../../files/extracted")

	require.Equal(archiveType, reader.ArchiveType(), "Archive Types didn't match %v - %v", archiveType, reader.ArchiveType())

	for {
		entry, err := reader.Next()
		if err == io.EOF {
			break
		}
		require.Nil(err, "Could not read entry from archive\n\t %v", err)
		if entry.IsDirectory() {
			continue
		}

		require.Equal(compressionType, entry.CompressionType())

		path := filepath.Join(tmp, entry.Name())
		WriteNewFile(path, 666)
	}

	err = CompareDirectories(extracted, tmp)
	require.Nil(err, "Directory compare failed %s %s\n\t %v", extracted, tmp, err)
}
