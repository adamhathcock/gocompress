package zip

import (
	"bytes"
	"os"
)

// IsZip checks the file has the Zip format signature by reading its beginning
// bytes and matching it against "PK\x03\x04"
func IsZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
