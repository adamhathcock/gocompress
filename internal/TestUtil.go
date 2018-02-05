package internal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/stretchr/testify/require"
)

type ReaderOpener func(archive string) (common.Reader, error)

func ExtractionTest(t *testing.T, archive string, opener ReaderOpener, archiveType common.ArchiveType, compressionType common.CompressionType) {
	require := require.New(t)

	aPath, err := filepath.Abs(archive)
	reader, err := opener(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)

	split := strings.Split(archive, "archives")
	extractedPath := split[0] + "extracted"

	tmp, err := MakeTempDir(".")

	require.Nil(err, "Could not create temp dir %v", err)

	tmp, err = filepath.Abs(tmp)
	defer os.RemoveAll(tmp)

	require.Nil(err, "Could not read abs extracted path\n\t %v", err)

	extracted, err := filepath.Abs(extractedPath)

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
