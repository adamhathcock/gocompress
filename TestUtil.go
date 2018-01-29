package gocompress

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExtractionTest(t *testing.T, reader Reader, archive string, archiveType ArchiveType) {
	require := require.New(t)

	tmp, err := MakeTempDir(".")

	require.Nil(err, "Could not create temp dir %v", err)

	tmp, err = filepath.Abs(tmp)
	defer os.RemoveAll(tmp)

	extracted, err := filepath.Abs("../../files/extracted")

	require.Nil(err, "Could not read abs extracted path\n\t %v", err)

	aPath, err := filepath.Abs("../../files/archives/" + archive)
	err = reader.OpenPath(aPath)

	require.Nil(err, "Could not open archive\n\t %v", err)

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

		path := filepath.Join(tmp, entry.Name())
		WriteNewFile(path, 666)
	}

	err = CompareDirectories(extracted, tmp)
	require.Nil(err, "Directory compare failed %s %s\n\t %v", extracted, tmp, err)
}
