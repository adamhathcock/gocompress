package tar

import (
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
)

func TestRarFormatReader_ReadEntry_Tar(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/tar/Tar.tar", OpenReader, common.TarArchive, common.None)
}

func TestRarFormatReader_ReadEntry_Tarbz2(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/tar/Tar.tar.bz2", OpenReader, common.TarArchive, common.BZip2)
}

func TestRarFormatReader_ReadEntry_Targz(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/tar/Tar.tar.gz", OpenReader, common.TarArchive, common.GZip)
}

func TestRarFormatReader_ReadEntry_Tarzx(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/tar/Tar.tar.xz", OpenReader, common.TarArchive, common.Xz)
}
