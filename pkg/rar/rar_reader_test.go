package rar

import (
	"testing"

	"github.com/adamhathcock/gocompress"
)

func TestRarFormatReader_ReadEntry_Rar(t *testing.T) {
	rr := &Reader
	gocompress.ExtractionTest(t, rr, "rar/Rar.rar", gocompress.RarArchive)
}
