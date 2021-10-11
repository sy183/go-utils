package csi

import "github.com/suy/utils/utils/flag"

type SgrCsi struct {
	srgFlags  flag.BitFlags
	fgColor8  uint8
	fgColor24 [3]uint8
	bgColor8  uint8
	bgColor24 [3]uint8
}

func NewSgrCsi() *SgrCsi {
	return &SgrCsi{srgFlags: flag.NewBitFlag(128)}
}

func (sc *SgrCsi) clearFgColorFlags() {
	sc.srgFlags[3] &= 0b00111111
	sc.srgFlags[4] = 0
	sc.srgFlags[14] &= 0b11111100
}

func (sc *SgrCsi) clearBgColorFlags() {
	sc.srgFlags[5] = 0
	sc.srgFlags[6] &= 0b11111100
	sc.srgFlags[14] &= 0b11110011
}

func (sc *SgrCsi) SetFgColor(fgColor flag.BitFlag) {
	sc.clearFgColorFlags()
	sc.srgFlags.SetFlag(fgColor)
}

func (sc *SgrCsi) SetBgColor(bgColor flag.BitFlag) {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(bgColor)
}

func (sc *SgrCsi) SetFgColor8(color8 uint8) {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomFgColorFlag)
	sc.srgFlags.SetFlag(FgColor8Flag)
	sc.fgColor8 = color8
}

func (sc *SgrCsi) SetBgColor8(color8 uint8) {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomBgColorFlag)
	sc.srgFlags.SetFlag(BgColor8Flag)
	sc.bgColor8 = color8
}
