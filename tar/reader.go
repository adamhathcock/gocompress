package tar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
	"github.com/dsnet/compress/bzip2"
	"github.com/ulikunitz/xz"
)

type Reader struct {
	tarReader   *tar.Reader
	compression common.CompressionType
}

type ReadCloser struct {
	Reader
	closer io.ReadCloser
}

func (readCloser *ReadCloser) Close() error {
	return readCloser.closer.Close()
}

func OpenReader(path string) (common.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
	}

	if isTarGzip(f) {
		f.Close()
		f, err = os.Open(path)
		gzip, err := gzip.NewReader(f)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		return open(gzip, common.GZip)
	}

	f.Close()
	f, err = os.Open(path)
	if isTarBz2(f) {
		f.Close()
		f, err = os.Open(path)
		bz2r, err := bzip2.NewReader(f, nil)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		return open(bz2r, common.BZip2)
	}

	f.Close()
	f, err = os.Open(path)
	if isTarXz(f) {
		f.Close()
		f, err = os.Open(path)
		xz, err := xz.NewReader(f)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		return open(&internal.ReadCloserWrapper{xz, f}, common.Xz)
	}

	f.Close()
	f, err = os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	return open(f, common.None)
}

func open(input io.ReadCloser, compressionType common.CompressionType) (common.ReadCloser, error) {
	var err error
	tfr := &ReadCloser{}
	tfr.tarReader = tar.NewReader(input)
	tfr.closer = input
	if tfr.tarReader == nil {
		return nil, fmt.Errorf("read: failed to create reader: %v", err)
	}
	tfr.compression = compressionType
	return tfr, nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (reader *Reader) Next() (common.Entry, error) {
	header, err := reader.tarReader.Next()
	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return &tarEntry{
		reader.tarReader,
		header,
		reader.compression}, nil
}

func (reader *Reader) ArchiveType() common.ArchiveType {
	return common.TarArchive
}
