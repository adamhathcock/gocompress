package zip

import (
	"testing"

	"github.com/adamhathcock/gocompress"
)

func TestZipFormatReader_ReadEntry_ZipDeflate(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "zip/Zip.zip", gocompress.ZipArchive, gocompress.Deflate)
}

func TestZipFormatReader_ReadEntry_Zip_Zip64Deflate(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "zip/Zip.zip64.zip", gocompress.ZipArchive, gocompress.Deflate)
}

func TestZipFormatReader_ReadEntry_Zip_Lzma(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "zip/Zip.Lzma.zip", gocompress.ZipArchive, gocompress.LZMA)
}

func TestZipFormatReader_ReadEntry_Zip_Deflate64(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "zip/Zip.deflate64.zip", gocompress.ZipArchive, gocompress.Deflate64)
}
