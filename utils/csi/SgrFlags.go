package csi

import (
	"bytes"
	"fmt"
)

type SgrFlags [16]byte

type SgrFlag int

const (
	ResetFlag      = 0 // 0
	BoldFlag       = 1
	FaintFlag      = 2
	ItalicFlag     = 3
	UnderlineFlag  = 4
	SlowBlinkFlag  = 5
	RapidBlinkFlag = 6
	ReversedFlag   = 7
	ConcealFlag    = 8 // 1
	CrossedOutFlag = 9

	PrimaryFontFlag      = 10
	AlternativeFont1Flag = 11
	AlternativeFont2Flag = 12
	AlternativeFont3Flag = 13
	AlternativeFont4Flag = 14
	AlternativeFont5Flag = 15
	AlternativeFont6Flag = 16 // 2
	AlternativeFont7Flag = 17
	AlternativeFont8Flag = 18
	AlternativeFont9Flag = 19

	FrakturFlag              = 20
	DoublyUnderlinedFlag     = 21
	NormalIntensityFlag      = 22
	NotItalicBlackLetterFlag = 23
	NotUnderlinedFlag        = 24 // 3
	NotBlinkingFlag          = 25
	ProportionalSpacingFlag  = 26
	NotReversedFlag          = 27
	RevealFlag               = 28
	NotCrossedOutFlag        = 29

	FgBlackFlag        = 30
	FgRedFlag          = 31
	FgGreenFlag        = 32 // 4
	FgYellowFlag       = 33
	FgBlueFlag         = 34
	FgMagentaFlag      = 35
	FgCyanFlag         = 36
	FgWhiteFlag        = 37
	CustomFgColorFlag  = 38
	DefaultFgColorFlag = 39

	BgBlackFlag        = 40 // 5
	BgRedFlag          = 41
	BgGreenFlag        = 42
	BgYellowFlag       = 43
	BgBlueFlag         = 44
	BgMagentaFlag      = 45
	BgCyanFlag         = 46
	BgWhiteFlag        = 47
	CustomBgColorFlag  = 48 // 6
	DefaultBgColorFlag = 49

	DisableProportionalSpacingFlag = 50
	FramedFlag                     = 51
	EncircledFlag                  = 52
	OverlinedFlag                  = 53
	NotFramedEncircledFlag         = 54
	NotOverlinedFlag               = 55
	SetUnderlineColorFlag          = 58 // 7
	DefaultUnderlineColorFlag      = 59

	IdeogramUnderlineFlag       = 60
	IdeogramDoubleUnderlineFlag = 61
	IdeogramOverlineFlag        = 62
	IdeogramDoubleOverlineFlag  = 63
	IdeogramStressMarkingFlag   = 64 // 8
	NoIdeogramAttributesFlag    = 65

	SuperscriptFlag             = 73 // 9
	SubscriptFlag               = 74
	NotSuperscriptSubscriptFlag = 75

	FgBrightBlackFlag   = 90 // 11
	FgBrightRedFlag     = 91
	FgBrightGreenFlag   = 92
	FgBrightYellowFlag  = 93
	FgBrightBlueFlag    = 94
	FgBrightMagentaFlag = 95
	FgBrightCyanFlag    = 96 // 12
	FgBrightWhiteFlag   = 97

	BgBrightBlackFlag   = 100
	BgBrightRedFlag     = 101
	BgBrightGreenFlag   = 102
	BgBrightYellowFlag  = 103
	BgBrightBlueFlag    = 104 // 13
	BgBrightMagentaFlag = 105
	BgBrightCyanFlag    = 106
	BgBrightWhiteFlag   = 107
)

func (sfs *SgrFlags) String() string {
	sb := bytes.NewBufferString("[")
	for i := 0; i < 16; i++ {
		if i == 15 {
			sb.WriteString(fmt.Sprintf("%08b]", sfs[i]))
		} else {
			sb.WriteString(fmt.Sprintf("%08b ", sfs[i]))
		}
	}
	return sb.String()
}

func (sfs *SgrFlags) SetSgrFlag(flag SgrFlag) {
	sfs[flag>>3] |= 1 << (flag & 7)
}

func (sfs *SgrFlags) UnsetSgrFlag(flag SgrFlag) {
	sfs[flag>>3] &= ^(1 << (flag & 7))
}

func (sfs *SgrFlags) SetSgrFlags(flags ...SgrFlag) {
	for _, flag := range flags {
		sfs[flag>>3] |= 1 << (flag & 7)
	}
}

func (sfs *SgrFlags) UnsetSgrFlags(flags ...SgrFlag) {
	for _, flag := range flags {
		sfs[flag>>3] &= ^(1 << (flag & 7))
	}
}

func (sfs *SgrFlags) UnsetAll() {
	*sfs = SgrFlags{}
}

func (sfs *SgrFlags) TestSgrFlag(flag SgrFlag) bool {
	return (sfs[flag>>3] & (1 << (flag & 7))) != 0
}

func (sfs *SgrFlags) TestSgrFlagUnset(flag SgrFlag) bool {
	return (sfs[flag>>3] & (1 << (flag & 7))) == 0
}

func (sfs *SgrFlags) TestAllSgrFlags(flags ...SgrFlag) bool {
	for _, flag := range flags {
		if (sfs[flag>>3] & (1 << (flag & 7))) == 0 {
			return false
		}
	}
	return true
}

func (sfs *SgrFlags) TestAllSgrFlagsUnset(flags ...SgrFlag) bool {
	for _, flag := range flags {
		if (sfs[flag>>3] & (1 << (flag & 7))) != 0 {
			return false
		}
	}
	return true
}

func (sfs *SgrFlags) TestHasSgrFlags(flags ...SgrFlag) bool {
	for _, flag := range flags {
		if (sfs[flag>>3] & (1 << (flag & 7))) != 0 {
			return true
		}
	}
	return false
}

func (sfs *SgrFlags) TestHasSgrFlagsUnset(flags ...SgrFlag) bool {
	for _, flag := range flags {
		if (sfs[flag>>3] & (1 << (flag & 7))) == 0 {
			return true
		}
	}
	return false
}

func (sfs *SgrFlags) clearFgColorFlags() {
	sfs[3] &= 0b00111111
	sfs[4] = 0
}

func (sfs *SgrFlags) clearBgColorFlags() {
	sfs[5] = 0
	sfs[6] &= 0b11111100
}

func (sfs *SgrFlags) SetFgColor(fgColor SgrFlag) {
	sfs.clearFgColorFlags()
	sfs.SetSgrFlag(fgColor)
}

func (sfs *SgrFlags) SetBgColor(bgColor SgrFlag) {
	sfs.clearBgColorFlags()
	sfs.SetSgrFlag(bgColor)
}

func (sfs *SgrFlags) SetCustomFgColorFlag() {
	sfs.SetSgrFlag(CustomFgColorFlag)
}

func (sfs *SgrFlags) SetCustomBgColorFlag() {
	sfs.SetSgrFlag(CustomBgColorFlag)
}
