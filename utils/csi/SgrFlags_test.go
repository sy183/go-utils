package csi

import (
	"fmt"
	"testing"
)

func TestSgrFlags(t *testing.T) {
	sc := NewSgrCsi()
	sc.srgFlags.SetFlags(BoldFlag, ItalicFlag, BgWhiteFlag)
	sc.fgColorRGB = ColorRGB{R: 134, G: 90, B: 234}
	sc.fgColor256 = 189
	fmt.Println(sc.String())
	fgColor := sc.getFgColor()
	fmt.Println(fgColor.fgString())
}
