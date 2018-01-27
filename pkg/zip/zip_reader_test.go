package zip

import (
	"testing"

	"github.com/adamhathcock/gocompress"
)


func TestRarFormatReader_ReadEntry_ZipDeflate(t *testing.T) {
	rr := &Reader
	gocompress.ExtractionTest(t, rr, "zip/Zip.zip")
}
