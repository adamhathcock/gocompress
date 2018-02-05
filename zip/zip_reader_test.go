package zip

import (
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
)

func TestZipFormatReader_ReadEntry_ZipDeflate(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/zip/Zip.zip", OpenReader, common.ZipArchive, common.Deflate)
}

func TestZipFormatReader_ReadEntry_Zip_Zip64Deflate(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/zip/Zip.zip64.zip", OpenReader, common.ZipArchive, common.Deflate)
}

func TestZipFormatReader_ReadEntry_Zip_Lzma(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/zip/Zip.Lzma.zip", OpenReader, common.ZipArchive, common.LZMA)
}

func TestZipFormatReader_ReadEntry_Zip_Deflate64(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/zip/Zip.deflate64.zip", OpenReader, common.ZipArchive, common.Deflate64)
}
