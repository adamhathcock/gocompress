package gocompress

// Reader is a Generic Archive Reader interface
type Reader interface {
	OpenPath(path string) error
	ReadEntry() (Entry, error)
	Close() error
}