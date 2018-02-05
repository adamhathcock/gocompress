package rar

import (
	"testing"

	"github.com/adamhathcock/gocompress/common"
	"github.com/adamhathcock/gocompress/internal"
)

func TestRarFormatReader_ReadEntry_Rar(t *testing.T) {
	internal.ExtractionTest(t, "../files/archives/rar/Rar.rar", OpenReader, common.RarArchive, common.Rar)
}
