package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress"
)

// Reader is the entry point for using an archive reader on a Rar archive
var Reader tarFormatReader

// IsTar uses the default tar implemented to check to see if the file is a tar
func IsTar(rarPath string) bool {
	f, err := os.Open(rarPath)
	if err != nil {
		return false
	}
	defer f.Close()

	reader := tar.NewReader(f)

	header, err := reader.Next()

	if err != nil {
		return false
	}

	switch header.Typeflag {
	case tar.TypeReg:
	case tar.TypeRegA:
	case tar.TypeLink:
	case tar.TypeSymlink:
	case tar.TypeChar:
	case tar.TypeBlock:
	case tar.TypeDir:
	case tar.TypeFifo:
	case tar.TypeCont:
	case tar.TypeXHeader:
	case tar.TypeXGlobalHeader:
	case tar.TypeGNULongName:
	case tar.TypeGNULongLink:
	case tar.TypeGNUSparse:
		return true
	}
	return false
}

type tarFormatReader struct {
	rarReader *tar.Reader
}

func (tfr *tarFormatReader) Close() error {
	return nil
}

func (tfr *tarFormatReader) OpenPath(path string) error {
	rf, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}

	return tfr.Open(rf)
}

func (tfr *tarFormatReader) Open(input io.Reader) error {
	var err error
	tfr.rarReader = tar.NewReader(input)
	if tfr.rarReader == nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (tfr *tarFormatReader) Next() (gocompress.Entry, error) {
	header, err := tfr.rarReader.Next()
	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	fmt.Println(string([]rune(header.Name)))
	return &tarFormatEntry{
		tfr.rarReader,
		header}, nil
}

func (tfr *tarFormatReader) ArchiveType() gocompress.ArchiveType {
	return gocompress.RarArchive
}
