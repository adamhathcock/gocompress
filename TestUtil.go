package gocompress

import (
	"os"
	"path/filepath"
	"testing"
)

func ExtractionTest(t *testing.T, reader Reader, archive string) {
	// if files come before their containing folders, then we must
	// create their folders before writing the file
	tmp, err := MakeTempDir(".")
	if err != nil {
		t.Fatalf("Could not create temp dir %v", err)
		return
	}
	tmp, err = filepath.Abs(tmp)
	defer os.RemoveAll(tmp)

	extracted, err := filepath.Abs("files/extracted")

	if err != nil {
		t.Fatalf("Could not read abs extracted path\n\t %v", err)
		return
	}

	aPath, err := filepath.Abs("files/archives/" + archive)
	err = reader.OpenPath(aPath)
	if err != nil {
		t.Fatalf("Could not open archive\n\t %v", err)
		return
	}
	for {
		entry, err := reader.ReadEntry()
		if err != nil {
			t.Fatalf("Could not read entry from archive\n\t %v", err)
			return
		}
		if entry == nil {
			break
		}
		if entry.IsDirectory() {
			continue
		}

		path := filepath.Join(tmp, entry.Name())
		WriteNewFile(path, 666)
	}

	err = CompareDirectories(extracted, tmp)
	if err != nil {
		t.Fatalf("Directory compare failed %s %s\n\t %v", extracted, tmp, err)
	}

}
