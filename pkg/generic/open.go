package generic

import (
	"github.com/adamhathcock/gocompress/pkg/rar"
	"github.com/adamhathcock/gocompress"
	"errors"
	"github.com/adamhathcock/gocompress/pkg/zip"
)

func OpenFilePath(path string) (gocompress.Reader, error) {
	if (rar.IsRar(path)) {
		rr := rar.Reader
		err := rr.OpenPath(path)
		if (err != nil) {
			return nil, err
		}
		// why do I have to do this?
		s := &rr
		return s, nil
	}
	if (zip.IsZip(path)) {
		zr := zip.Reader
		err := zr.OpenPath(path)
		if (err != nil) {
			return nil, err
		}
		return &zr, nil
	}
	return nil, errors.New(path + " has no valid format detected")
}
