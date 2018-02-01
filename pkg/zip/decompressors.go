package zip

import (
	"compress/bzip2"
	"io"
)

type BZip2Reader struct {
	reader io.Reader
}

func (reader BZip2Reader) Read(p []byte) (n int, err error) {
	return reader.reader.Read(p)
}

func (reader BZip2Reader) Close() error {
	return nil
}

func BZip2Decompressor(r io.Reader) io.ReadCloser {
	return BZip2Reader{bzip2.NewReader(r)}
}
