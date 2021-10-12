package csi

import (
	"encoding/json"
	"github.com/suy/utils/utils/flag"
	"github.com/suy/utils/utils/uns"
)

type SgrCsi struct {
	srgFlags   flag.BitFlags
	fgColor256 Color256
	fgColorRGB ColorRGB
	bgColor256 Color256
	bgColorRGB ColorRGB
}

func NewSgrCsi() *SgrCsi {
	return &SgrCsi{srgFlags: flag.NewBitFlag(128)}
}

func (sc *SgrCsi) String() string {
	m := map[string]interface{}{
		"Reset":      sc.srgFlags.TestFlag(ResetFlag),
		"Bold":       sc.srgFlags.TestFlag(BoldFlag),
		"Faint":      sc.srgFlags.TestFlag(FaintFlag),
		"Italic":     sc.srgFlags.TestFlag(ItalicFlag),
		"Underline":  sc.srgFlags.TestFlag(UnderlineFlag),
		"SlowBlink":  sc.srgFlags.TestFlag(SlowBlinkFlag),
		"RapidBlink": sc.srgFlags.TestFlag(RapidBlinkFlag),
		"Reversed":   sc.srgFlags.TestFlag(ReversedFlag),
		"Conceal":    sc.srgFlags.TestFlag(ConcealFlag),
		"CrossedOut": sc.srgFlags.TestFlag(CrossedOutFlag),

		"PrimaryFont":      sc.srgFlags.TestFlag(PrimaryFontFlag),
		"AlternativeFont1": sc.srgFlags.TestFlag(AlternativeFont1Flag),
		"AlternativeFont2": sc.srgFlags.TestFlag(AlternativeFont2Flag),
		"AlternativeFont3": sc.srgFlags.TestFlag(AlternativeFont3Flag),
		"AlternativeFont4": sc.srgFlags.TestFlag(AlternativeFont4Flag),
		"AlternativeFont5": sc.srgFlags.TestFlag(AlternativeFont5Flag),
		"AlternativeFont6": sc.srgFlags.TestFlag(AlternativeFont6Flag),
		"AlternativeFont7": sc.srgFlags.TestFlag(AlternativeFont7Flag),
		"AlternativeFont8": sc.srgFlags.TestFlag(AlternativeFont8Flag),
		"AlternativeFont9": sc.srgFlags.TestFlag(AlternativeFont9Flag),

		"Fraktur":              sc.srgFlags.TestFlag(FrakturFlag),
		"DoublyUnderlined":     sc.srgFlags.TestFlag(DoublyUnderlinedFlag),
		"NormalIntensity":      sc.srgFlags.TestFlag(NormalIntensityFlag),
		"NotItalicBlackLetter": sc.srgFlags.TestFlag(NotItalicBlackLetterFlag),
		"NotUnderlined":        sc.srgFlags.TestFlag(NotUnderlinedFlag),
		"NotBlinking":          sc.srgFlags.TestFlag(NotBlinkingFlag),
		"ProportionalSpacing":  sc.srgFlags.TestFlag(ProportionalSpacingFlag),
		"NotReversed":          sc.srgFlags.TestFlag(NotReversedFlag),
		"Reveal":               sc.srgFlags.TestFlag(RevealFlag),
		"NotCrossedOut":        sc.srgFlags.TestFlag(NotCrossedOutFlag),

		"FgBlack":        sc.srgFlags.TestFlag(FgBlackFlag),
		"FgRed":          sc.srgFlags.TestFlag(FgRedFlag),
		"FgGreen":        sc.srgFlags.TestFlag(FgGreenFlag),
		"FgYellow":       sc.srgFlags.TestFlag(FgYellowFlag),
		"FgBlue":         sc.srgFlags.TestFlag(FgBlueFlag),
		"FgMagenta":      sc.srgFlags.TestFlag(FgMagentaFlag),
		"FgCyan":         sc.srgFlags.TestFlag(FgCyanFlag),
		"FgWhite":        sc.srgFlags.TestFlag(FgWhiteFlag),
		"CustomFgColor":  sc.srgFlags.TestFlag(CustomFgColorFlag),
		"DefaultFgColor": sc.srgFlags.TestFlag(DefaultFgColorFlag),

		"BgBlack":        sc.srgFlags.TestFlag(BgBlackFlag),
		"BgRed":          sc.srgFlags.TestFlag(BgRedFlag),
		"BgGreen":        sc.srgFlags.TestFlag(BgGreenFlag),
		"BgYellow":       sc.srgFlags.TestFlag(BgYellowFlag),
		"BgBlue":         sc.srgFlags.TestFlag(BgBlueFlag),
		"BgMagenta":      sc.srgFlags.TestFlag(BgMagentaFlag),
		"BgCyan":         sc.srgFlags.TestFlag(BgCyanFlag),
		"BgWhite":        sc.srgFlags.TestFlag(BgWhiteFlag),
		"CustomBgColor":  sc.srgFlags.TestFlag(CustomBgColorFlag),
		"DefaultBgColor": sc.srgFlags.TestFlag(DefaultBgColorFlag),

		"DisableProportionalSpacing": sc.srgFlags.TestFlag(DisableProportionalSpacingFlag),
		"Framed":                     sc.srgFlags.TestFlag(FramedFlag),
		"Encircled":                  sc.srgFlags.TestFlag(EncircledFlag),
		"Overlined":                  sc.srgFlags.TestFlag(OverlinedFlag),
		"NotFramedEncircled":         sc.srgFlags.TestFlag(NotFramedEncircledFlag),
		"NotOverlined":               sc.srgFlags.TestFlag(NotOverlinedFlag),
		"SetUnderlineColor":          sc.srgFlags.TestFlag(SetUnderlineColorFlag),
		"DefaultUnderlineColor":      sc.srgFlags.TestFlag(DefaultUnderlineColorFlag),

		"IdeogramUnderline":       sc.srgFlags.TestFlag(IdeogramUnderlineFlag),
		"IdeogramDoubleUnderline": sc.srgFlags.TestFlag(IdeogramDoubleUnderlineFlag),
		"IdeogramOverline":        sc.srgFlags.TestFlag(IdeogramOverlineFlag),
		"IdeogramDoubleOverline":  sc.srgFlags.TestFlag(IdeogramDoubleOverlineFlag),
		"IdeogramStressMarking":   sc.srgFlags.TestFlag(IdeogramStressMarkingFlag),
		"NoIdeogramAttributes":    sc.srgFlags.TestFlag(NoIdeogramAttributesFlag),

		"Superscript":             sc.srgFlags.TestFlag(SuperscriptFlag),
		"Subscript":               sc.srgFlags.TestFlag(SubscriptFlag),
		"NotSuperscriptSubscript": sc.srgFlags.TestFlag(NotSuperscriptSubscriptFlag),

		"FgBrightBlack":   sc.srgFlags.TestFlag(FgBrightBlackFlag),
		"FgBrightRed":     sc.srgFlags.TestFlag(FgBrightRedFlag),
		"FgBrightGreen":   sc.srgFlags.TestFlag(FgBrightGreenFlag),
		"FgBrightYellow":  sc.srgFlags.TestFlag(FgBrightYellowFlag),
		"FgBrightBlue":    sc.srgFlags.TestFlag(FgBrightBlueFlag),
		"FgBrightMagenta": sc.srgFlags.TestFlag(FgBrightMagentaFlag),
		"FgBrightCyan":    sc.srgFlags.TestFlag(FgBrightCyanFlag),
		"FgBrightWhite":   sc.srgFlags.TestFlag(FgBrightWhiteFlag),

		"BgBrightBlack":   sc.srgFlags.TestFlag(BgBrightBlackFlag),
		"BgBrightRed":     sc.srgFlags.TestFlag(BgBrightRedFlag),
		"BgBrightGreen":   sc.srgFlags.TestFlag(BgBrightGreenFlag),
		"BgBrightYellow":  sc.srgFlags.TestFlag(BgBrightYellowFlag),
		"BgBrightBlue":    sc.srgFlags.TestFlag(BgBrightBlueFlag),
		"BgBrightMagenta": sc.srgFlags.TestFlag(BgBrightMagentaFlag),
		"BgBrightCyan":    sc.srgFlags.TestFlag(BgBrightCyanFlag),
		"BgBrightWhite":   sc.srgFlags.TestFlag(BgBrightWhiteFlag),

		"FgColor256": sc.srgFlags.TestFlag(FgColor256Flag),
		"FgColorRGB": sc.srgFlags.TestFlag(FgColorRGBFlag),
		"BgColor256": sc.srgFlags.TestFlag(BgColor256Flag),
		"BgColorRGB": sc.srgFlags.TestFlag(BgColorRGBFlag),

		"fgColor256": sc.fgColor256,
		"fgColorRGB": map[string]uint8{
			"R": sc.fgColorRGB.R,
			"G": sc.fgColorRGB.G,
			"B": sc.fgColorRGB.B,
		},
		"bgColor256": sc.bgColor256,
		"bgColorRGB": map[string]uint8{
			"R": sc.bgColorRGB.R,
			"G": sc.bgColorRGB.G,
			"B": sc.bgColorRGB.B,
		},
	}
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err.Error()
	}
	return uns.BytesToString(bytes)
}

func (sc *SgrCsi) getFgColor() *Color {
	return &Color{
		normalFlags: (sc.srgFlags[3] >> 6) | (sc.srgFlags[4] << 2),
		brightFlags: (sc.srgFlags[11] >> 2) | (sc.srgFlags[12] << 6),
		customFlag:  sc.srgFlags.TestFlag(CustomFgColorFlag),
		defaultFlag: sc.srgFlags.TestFlag(DefaultFgColorFlag),
		rgb256Flags: sc.srgFlags[14] & 0b00000011,
		color256:    sc.fgColor256,
		colorRGB:    sc.fgColorRGB,
	}
}

func (sc *SgrCsi) getBgColor() *Color {
	return &Color{
		normalFlags: sc.srgFlags[5],
		brightFlags: (sc.srgFlags[12] >> 4) | (sc.srgFlags[13] << 4),
		customFlag:  sc.srgFlags.TestFlag(CustomBgColorFlag),
		defaultFlag: sc.srgFlags.TestFlag(DefaultBgColorFlag),
		rgb256Flags: (sc.srgFlags[14] & 0b00001100) >> 2,
		color256:    sc.bgColor256,
		colorRGB:    sc.bgColorRGB,
	}
}

func (sc *SgrCsi) clearFgColorFlags() {
	sc.srgFlags[3] &= 0b00111111
	sc.srgFlags[4] = 0
	sc.srgFlags[11] &= 0b00000011
	sc.srgFlags[12] &= 0b11111100
	sc.srgFlags[14] &= 0b11111100
	sc.fgColor256 = 0
	sc.fgColorRGB = ColorRGB{}
}

func (sc *SgrCsi) clearBgColorFlags() {
	sc.srgFlags[5] = 0
	sc.srgFlags[6] &= 0b11111100
	sc.srgFlags[12] &= 0b00001111
	sc.srgFlags[13] &= 0b11110000
	sc.srgFlags[14] &= 0b11110011
	sc.bgColor256 = 0
	sc.bgColorRGB = ColorRGB{}
}

func (sc *SgrCsi) SetFgColor(fgColor flag.BitFlag) *SgrCsi {
	sc.clearFgColorFlags()
	sc.srgFlags.SetFlag(fgColor)
	return sc
}

func (sc *SgrCsi) SetBgColor(bgColor flag.BitFlag) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(bgColor)
	return sc
}

func (sc *SgrCsi) SetFgColor256(fgColor256 Color256) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomFgColorFlag)
	sc.srgFlags.SetFlag(FgColor256Flag)
	sc.fgColor256 = fgColor256
	return sc
}

func (sc *SgrCsi) SetBgColor256(bgColor256 Color256) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomBgColorFlag)
	sc.srgFlags.SetFlag(BgColor256Flag)
	sc.bgColor256 = bgColor256
	return sc
}

func (sc *SgrCsi) SetFgColorRGB(fgColorRGB ColorRGB) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomFgColorFlag)
	sc.srgFlags.SetFlag(FgColor256Flag)
	sc.fgColorRGB = fgColorRGB
	return sc
}

func (sc *SgrCsi) SetBgColorRGB(bgColorRGB ColorRGB) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(CustomBgColorFlag)
	sc.srgFlags.SetFlag(BgColor256Flag)
	sc.bgColorRGB = bgColorRGB
	return sc
}
