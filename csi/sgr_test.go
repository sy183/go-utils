package csi

import (
	"fmt"
	"testing"
)

func TestSgrFlags(t *testing.T) {
	//sc := NewSgrCsi()
	//sc.srgFlags.SetFlags(BoldFlag, ItalicFlag, CrossedOutFlag, CustomFgColorFlag, FgColor256Flag)
	//sc.fgColorRGB = ColorRGB{R: 255, G: 128, B: 0}
	//sc.fgColor256 = 121
	//scs := sc.String()
	//fmt.Println("==================")
	//fmt.Println(scs + "hello" + NewSgrCsi().Reset().String())
	for i := 0; i < 256; i++ {
		fmt.Println(NewSgrCsi().SetFgColor256(Color256(i)).String() + "hello" + NewSgrCsi().Reset().String())
	}
}
