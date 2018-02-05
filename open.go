package gocompress

import (
	"errors"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/rar"
	"github.com/adamhathcock/gocompress/tar"
	"github.com/adamhathcock/gocompress/zip"
)

func OpenReader(path string) (common.ReadCloser, error) {
	if rar.IsRar(path) {
		reader, err := rar.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	if zip.IsZip(path) {
		reader, err := zip.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	if tar.IsTar(path) {
		reader, err := tar.OpenReader(path)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	return nil, errors.New(path + " has no valid format detected")
}
