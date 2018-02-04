package tar

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/adamhathcock/gocompress"
	"github.com/dsnet/compress/bzip2"
	"github.com/ulikunitz/xz"
)

const tarBlockSize int = 512

// IsTar uses the default tar implemented to check to see if the file is a tar
func IsTar(rarPath string) bool {
	f, err := os.Open(rarPath)
	if err != nil {
		return false
	}

	if isTarBz2(f) {
		return true
	}
	f.Close()
	f, err = os.Open(rarPath)
	if err != nil {
		return false
	}
	defer f.Close()
	return isTar(f)
}

func isTar(r io.Reader) bool {

	reader := tar.NewReader(r)

	header, err := reader.Next()

	if err != nil {
		return false
	}

	switch header.Typeflag {
	case tar.TypeReg, tar.TypeRegA, tar.TypeLink, tar.TypeSymlink, tar.TypeChar,
		tar.TypeBlock, tar.TypeDir, tar.TypeFifo, tar.TypeCont, tar.TypeXHeader,
		tar.TypeXGlobalHeader, tar.TypeGNULongName, tar.TypeGNULongLink, tar.TypeGNUSparse:
		return true
	default:
		return false
	}
	return false
}

// hasTarHeader checks passed bytes has a valid tar header or not. buf must
// contain at least 512 bytes and if not, it always returns false.
func hasTarHeader(buf []byte) bool {
	if len(buf) < tarBlockSize {
		return false
	}

	b := buf[148:156]
	b = bytes.Trim(b, " \x00") // clean up all spaces and null bytes
	if len(b) == 0 {
		return false // unknown format
	}
	hdrSum, err := strconv.ParseUint(string(b), 8, 64)
	if err != nil {
		return false
	}

	// According to the go official archive/tar, Sun tar uses signed byte
	// values so this calcs both signed and unsigned
	var usum uint64
	var sum int64
	for i, c := range buf {
		if 148 <= i && i < 156 {
			c = ' ' // checksum field itself is counted as branks
		}
		usum += uint64(uint8(c))
		sum += int64(int8(c))
	}

	if hdrSum != usum && int64(hdrSum) != sum {
		return false // invalid checksum
	}

	return true
}

func isTarXz(f io.Reader) bool {
	xz, err := xz.NewReader(f)
	if err != nil {
		return false
	}
	buf := make([]byte, tarBlockSize)
	n, err := xz.Read(buf)
	if err != nil || n < tarBlockSize {
		return false
	}

	return hasTarHeader(buf)
}

func isTarGzip(f io.Reader) bool {
	gzip, err := gzip.NewReader(f)
	if err != nil {
		return false
	}
	defer gzip.Close()
	buf := make([]byte, tarBlockSize)
	n, err := gzip.Read(buf)
	if err != nil || n < tarBlockSize {
		return false
	}

	return hasTarHeader(buf)
}

func isTarBz2(f io.Reader) bool {
	bz2r, err := bzip2.NewReader(f, nil)
	if err != nil {
		return false
	}
	defer bz2r.Close()
	buf := make([]byte, tarBlockSize)
	n, err := bz2r.Read(buf)
	if err != nil || n < tarBlockSize {
		return false
	}

	return hasTarHeader(buf)
}

type Reader struct {
	rarReader   *tar.Reader
	compression gocompress.CompressionType
}

func (tfr *Reader) Close() error {
	return nil
}

func (tfr *Reader) OpenPath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}

	if isTarGzip(f) {
		f.Close()
		f, err = os.Open(path)
		gzip, err := gzip.NewReader(f)
		if err != nil {
			return fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		tfr.compression = gocompress.GZip
		return tfr.Open(gzip)
	}

	f.Close()
	f, err = os.Open(path)
	if isTarBz2(f) {
		f.Close()
		f, err = os.Open(path)
		bz2r, err := bzip2.NewReader(f, nil)
		if err != nil {
			return fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		tfr.compression = gocompress.BZip2
		return tfr.Open(bz2r)
	}

	f.Close()
	f, err = os.Open(path)
	if isTarXz(f) {
		f.Close()
		f, err = os.Open(path)
		bz2r, err := xz.NewReader(f)
		if err != nil {
			return fmt.Errorf("%s: failed to open file: %v", path, err)
		}
		tfr.compression = gocompress.Xz
		return tfr.Open(bz2r)
	}

	f.Close()
	f, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: failed to open file: %v", path, err)
	}
	tfr.compression = gocompress.None
	return tfr.Open(f)
}

func (tfr *Reader) Open(input io.Reader) error {
	var err error
	tfr.rarReader = tar.NewReader(input)
	if tfr.rarReader == nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	return nil
}

// Read extracts the RAR file read from input and puts the contents
// into destination.
func (tfr *Reader) Next() (gocompress.Entry, error) {
	header, err := tfr.rarReader.Next()
	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return &tarFormatEntry{
		tfr.rarReader,
		header,
		tfr.compression}, nil
}

func (tfr *Reader) ArchiveType() gocompress.ArchiveType {
	return gocompress.TarArchive
}
