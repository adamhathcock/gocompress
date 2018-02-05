package internal

import (
	"io"

)

type ReadCloserWrapper struct {
	Reader io.Reader
	Closer io.ReadCloser
}


func (wrapper *ReadCloserWrapper) Close() error {
	return wrapper.Closer.Close()
}

func (wrapper *ReadCloserWrapper) Read(p []byte) (n int, err error) {
	return wrapper.Reader.Read(p)
}