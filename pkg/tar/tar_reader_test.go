package tar

import (
	"testing"

	"github.com/adamhathcock/gocompress"
)

func TestRarFormatReader_ReadEntry_Tar(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "tar/Tar.tar", gocompress.TarArchive, gocompress.None)
}

func TestRarFormatReader_ReadEntry_Tarbz2(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "tar/Tar.tar.bz2", gocompress.TarArchive, gocompress.BZip2)
}

func TestRarFormatReader_ReadEntry_Targz(t *testing.T) {
	rr := &Reader{}
	gocompress.ExtractionTest(t, rr, "tar/Tar.tar.gz", gocompress.TarArchive, gocompress.GZip)
}