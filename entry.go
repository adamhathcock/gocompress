package gocompress

import (
	"errors"
	"io"
	"os"
)

// Entry is the generic archive entry interface when reading archives
type Entry interface {
	Name() string
	IsDirectory() bool
	Mode() os.FileMode
	Write(output io.Writer) error
}

type nilEntry struct {
}

func (entry nilEntry) Name() string {
	return "nil"
}

func (entry nilEntry) IsDirectory() bool {
	return false
}

func (entry nilEntry) Mode() os.FileMode {
	return os.ModeAppend
}

func (entry nilEntry) Write(output io.Writer) error {
	return errors.New("nil")
}

// NilEntry is the null entry when dealing with Entry objects
var NilEntry = nilEntry{}
