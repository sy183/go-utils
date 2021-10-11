package csi

import (
	"fmt"
	"testing"
)

func TestSgrFlags(t *testing.T) {
	sfs := &SgrFlags{}
	sfs.SetSgrFlags(BoldFlag, ItalicFlag, FgBlackFlag, BgWhiteFlag)
	fmt.Println(sfs)
	sfs.clearFgColorFlags()
	fmt.Println(sfs)
}
