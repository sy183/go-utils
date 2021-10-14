package csi

import (
	"bytes"
	"github.com/suy/utils/flag"
	"github.com/suy/utils/uns"
	"strconv"
)

const (
	resetSgrCsi = "\x1b[0m"
	sgrCsiHead  = "\x1b["
	sgrCsiEnd   = 'm'
)

var uint8ToStringTable [256]string

func initUint8ToStringTable() {
	for i := 0; i < 256; i++ {
		uint8ToStringTable[i] = strconv.Itoa(i)
	}
}

func init() {
	initUint8ToStringTable()
}

type SgrCsi struct {
	srgFlags          flag.BitFlags
	fgColor256        Color256
	fgColorRGB        ColorRGB
	bgColor256        Color256
	bgColorRGB        ColorRGB
	underlineColor256 Color256
	underlineColorRGB ColorRGB
	cache             string
	cacheValid        bool
}

func NewSgrCsi() *SgrCsi {
	return &SgrCsi{srgFlags: flag.NewBitFlag(128), cacheValid: true}
}

func (sc *SgrCsi) Clean() *SgrCsi {
	*sc = SgrCsi{srgFlags: sc.srgFlags}
	sc.srgFlags.Clean()
	sc.cacheValid = true
	return sc
}

func (sc *SgrCsi) Reset() *SgrCsi {
	sc.Clean()
	sc.srgFlags.SetFlag(ResetFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) IsReset() bool {
	return sc.srgFlags.TestFlag(ResetFlag)
}

func (sc *SgrCsi) SetBold() *SgrCsi {
	sc.srgFlags.UnsetFlags(NormalIntensityFlag, FaintFlag)
	sc.srgFlags.SetFlag(BoldFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetFaint() *SgrCsi {
	sc.srgFlags.UnsetFlags(NormalIntensityFlag, BoldFlag)
	sc.srgFlags.SetFlag(FaintFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNormalIntensity() *SgrCsi {
	sc.srgFlags.UnsetFlags(BoldFlag, FaintFlag)
	sc.srgFlags.SetFlag(NormalIntensityFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetItalic() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotItalicBlackLetterFlag, FrakturFlag)
	sc.srgFlags.SetFlag(ItalicFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetFraktur() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotItalicBlackLetterFlag, ItalicFlag)
	sc.srgFlags.SetFlag(FrakturFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotItalicBlackLetter() *SgrCsi {
	sc.srgFlags.UnsetFlags(ItalicFlag, FrakturFlag)
	sc.srgFlags.SetFlag(NotItalicBlackLetterFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetReversed() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotReversedFlag)
	sc.srgFlags.SetFlag(ReversedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotReversed() *SgrCsi {
	sc.srgFlags.UnsetFlags(ReversedFlag)
	sc.srgFlags.SetFlag(NotReversedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetUnderlined() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotUnderlinedFlag, DoublyUnderlinedFlag)
	sc.srgFlags.SetFlag(UnderlineFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetDoublyUnderlined() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotUnderlinedFlag, UnderlineFlag)
	sc.srgFlags.SetFlag(DoublyUnderlinedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotUnderlined() *SgrCsi {
	sc.srgFlags.UnsetFlags(UnderlineFlag, DoublyUnderlinedFlag)
	sc.srgFlags.SetFlag(NotUnderlinedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetSlowBlink() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotBlinkingFlag, RapidBlinkFlag)
	sc.srgFlags.SetFlag(SlowBlinkFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetRapidBlink() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotBlinkingFlag, SlowBlinkFlag)
	sc.srgFlags.SetFlag(RapidBlinkFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotBlinking() *SgrCsi {
	sc.srgFlags.UnsetFlags(SlowBlinkFlag, RapidBlinkFlag)
	sc.srgFlags.SetFlag(NotBlinkingFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetOverlined() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotOverlinedFlag)
	sc.srgFlags.SetFlag(OverlinedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotOverlined() *SgrCsi {
	sc.srgFlags.UnsetFlags(OverlinedFlag)
	sc.srgFlags.SetFlag(NotOverlinedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetConceal() *SgrCsi {
	sc.srgFlags.UnsetFlags(RevealFlag)
	sc.srgFlags.SetFlag(ConcealFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetReveal() *SgrCsi {
	sc.srgFlags.UnsetFlags(ConcealFlag)
	sc.srgFlags.SetFlag(RevealFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetCrossedOut() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotCrossedOutFlag)
	sc.srgFlags.SetFlag(CrossedOutFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotCrossedOut() *SgrCsi {
	sc.srgFlags.UnsetFlags(CrossedOutFlag)
	sc.srgFlags.SetFlag(NotCrossedOutFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetProportionalSpacing() *SgrCsi {
	sc.srgFlags.UnsetFlags(DisableProportionalSpacingFlag)
	sc.srgFlags.SetFlag(ProportionalSpacingFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetDisableProportionalSpacing() *SgrCsi {
	sc.srgFlags.UnsetFlags(ProportionalSpacingFlag)
	sc.srgFlags.SetFlag(DisableProportionalSpacingFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetFramed() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotFramedEncircledFlag, EncircledFlag)
	sc.srgFlags.SetFlag(FramedFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetEncircled() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotFramedEncircledFlag, FramedFlag)
	sc.srgFlags.SetFlag(EncircledFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotFramedEncircled() *SgrCsi {
	sc.srgFlags.UnsetFlags(FramedFlag, EncircledFlag)
	sc.srgFlags.SetFlag(NotFramedEncircledFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetSuperscript() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotSuperscriptSubscriptFlag, SubscriptFlag)
	sc.srgFlags.SetFlag(SuperscriptFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetSubscript() *SgrCsi {
	sc.srgFlags.UnsetFlags(NotSuperscriptSubscriptFlag, SuperscriptFlag)
	sc.srgFlags.SetFlag(SubscriptFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNotSuperscriptSubscript() *SgrCsi {
	sc.srgFlags.UnsetFlags(SuperscriptFlag, SubscriptFlag)
	sc.srgFlags.SetFlag(NotSuperscriptSubscriptFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetIdeogramUnderline() *SgrCsi {
	sc.srgFlags[5] &= 0b11000001
	sc.srgFlags.SetFlag(IdeogramUnderlineFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetIdeogramDoubleUnderline() *SgrCsi {
	sc.srgFlags[5] &= 0b11000010
	sc.srgFlags.SetFlag(IdeogramDoubleUnderlineFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetIdeogramOverline() *SgrCsi {
	sc.srgFlags[5] &= 0b11000100
	sc.srgFlags.SetFlag(IdeogramOverlineFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetIdeogramDoubleOverline() *SgrCsi {
	sc.srgFlags[5] &= 0b11001000
	sc.srgFlags.SetFlag(IdeogramDoubleOverlineFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetIdeogramStressMarking() *SgrCsi {
	sc.srgFlags[5] &= 0b11010000
	sc.srgFlags.SetFlag(IdeogramStressMarkingFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetNoIdeogramAttributes() *SgrCsi {
	sc.srgFlags[5] &= 0b11100000
	sc.srgFlags.SetFlag(NoIdeogramAttributesFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) IsBold() bool {
	 return sc.srgFlags.TestFlag(BoldFlag)
}

func (sc *SgrCsi) IsFaint() bool {
	 return sc.srgFlags.TestFlag(FaintFlag)
}

func (sc *SgrCsi) IsNormalIntensity() bool {
	 return sc.srgFlags.TestFlag(NormalIntensityFlag)
}

func (sc *SgrCsi) IsItalic() bool {
	 return sc.srgFlags.TestFlag(ItalicFlag)
}

func (sc *SgrCsi) IsFraktur() bool {
	 return sc.srgFlags.TestFlag(FrakturFlag)
}

func (sc *SgrCsi) IsNotItalicBlackLetter() bool {
	 return sc.srgFlags.TestFlag(NotItalicBlackLetterFlag)
}

func (sc *SgrCsi) IsReversed() bool {
	 return sc.srgFlags.TestFlag(ReversedFlag)
}

func (sc *SgrCsi) IsNotReversed() bool {
	 return sc.srgFlags.TestFlag(NotReversedFlag)
}

func (sc *SgrCsi) IsUnderlined() bool {
	 return sc.srgFlags.TestFlag(UnderlineFlag)
}

func (sc *SgrCsi) IsDoublyUnderlined() bool {
	 return sc.srgFlags.TestFlag(DoublyUnderlinedFlag)
}

func (sc *SgrCsi) NotUnderlined() bool {
	 return sc.srgFlags.TestFlag(NotUnderlinedFlag)
}

func (sc *SgrCsi) IsSlowBlink() bool {
	 return sc.srgFlags.TestFlag(SlowBlinkFlag)
}

func (sc *SgrCsi) IsRapidBlink() bool {
	 return sc.srgFlags.TestFlag(RapidBlinkFlag)
}

func (sc *SgrCsi) NotBlinking() bool {
	 return sc.srgFlags.TestFlag(NotBlinkingFlag)
}

func (sc *SgrCsi) IsOverlined() bool {
	 return sc.srgFlags.TestFlag(OverlinedFlag)
}

func (sc *SgrCsi) NotOverlined() bool {
	 return sc.srgFlags.TestFlag(NotOverlinedFlag)
}

func (sc *SgrCsi) IsConceal() bool {
	 return sc.srgFlags.TestFlag(ConcealFlag)
}

func (sc *SgrCsi) IsReveal() bool {
	 return sc.srgFlags.TestFlag(RevealFlag)
}

func (sc *SgrCsi) IsCrossedOut() bool {
	 return sc.srgFlags.TestFlag(CrossedOutFlag)
}

func (sc *SgrCsi) NotCrossedOut() bool {
	 return sc.srgFlags.TestFlag(NotCrossedOutFlag)
}

func (sc *SgrCsi) IsProportionalSpacing() bool {
	 return sc.srgFlags.TestFlag(ProportionalSpacingFlag)
}

func (sc *SgrCsi) DisabledProportionalSpacing() bool {
	 return sc.srgFlags.TestFlag(DisableProportionalSpacingFlag)
}

func (sc *SgrCsi) IsFramed() bool {
	 return sc.srgFlags.TestFlag(FramedFlag)
}

func (sc *SgrCsi) IsEncircled() bool {
	 return sc.srgFlags.TestFlag(EncircledFlag)
}

func (sc *SgrCsi) NotFramedEncircled() bool {
	 return sc.srgFlags.TestFlag(NotFramedEncircledFlag)
}

func (sc *SgrCsi) IsSuperscript() bool {
	 return sc.srgFlags.TestFlag(SuperscriptFlag)
}

func (sc *SgrCsi) IsSubscript() bool {
	 return sc.srgFlags.TestFlag(SubscriptFlag)
}

func (sc *SgrCsi) NotSuperscriptSubscript() bool {
	 return sc.srgFlags.TestFlag(NotSuperscriptSubscriptFlag)
}

func (sc *SgrCsi) IsIdeogramUnderline() bool {
	 return sc.srgFlags.TestFlag(IdeogramUnderlineFlag)
}

func (sc *SgrCsi) IsIdeogramDoubleUnderline() bool {
	 return sc.srgFlags.TestFlag(IdeogramDoubleUnderlineFlag)
}

func (sc *SgrCsi) IsIdeogramOverline() bool {
	 return sc.srgFlags.TestFlag(IdeogramOverlineFlag)
}

func (sc *SgrCsi) IsIdeogramDoubleOverline() bool {
	 return sc.srgFlags.TestFlag(IdeogramDoubleOverlineFlag)
}

func (sc *SgrCsi) IsIdeogramStressMarking() bool {
	 return sc.srgFlags.TestFlag(IdeogramStressMarkingFlag)
}

func (sc *SgrCsi) NotIdeogramAttributes() bool {
	 return sc.srgFlags.TestFlag(NoIdeogramAttributesFlag)
}

func (sc *SgrCsi) clearFgColorFlags() {
	sc.srgFlags[6] = 0
	sc.srgFlags[8] = 0
	sc.srgFlags[10] &= 0b11111100
	sc.srgFlags[11] &= 0b11111100
	sc.fgColor256 = 0
	sc.fgColorRGB = ColorRGB{}
}

func (sc *SgrCsi) clearBgColorFlags() {
	sc.srgFlags[7] = 0
	sc.srgFlags[9] = 0
	sc.srgFlags[10] &= 0b11110011
	sc.srgFlags[11] &= 0b11110011
	sc.bgColor256 = 0
	sc.bgColorRGB = ColorRGB{}
}

func (sc *SgrCsi) clearUnderlineColorFlags() {
	sc.srgFlags[10] &= 0b11001111
	sc.srgFlags[11] &= 0b11001111
	sc.underlineColor256 = 0
	sc.underlineColorRGB = ColorRGB{}
}

func (sc *SgrCsi) SetFgColor(fgColor flag.BitFlag) *SgrCsi {
	if (fgColor >= FgBlackFlag && fgColor <= FgWhiteFlag) || (fgColor >= FgBrightBlackFlag && fgColor <= FgBrightWhiteFlag) {
		sc.clearFgColorFlags()
		sc.srgFlags.SetFlag(fgColor)
		sc.cacheValid = false
	}
	return sc
}

func (sc *SgrCsi) SetBgColor(bgColor flag.BitFlag) *SgrCsi {
	if (bgColor >= BgBlackFlag && bgColor <= BgWhiteFlag) || (bgColor >= BgBrightBlackFlag && bgColor <= BgBrightWhiteFlag) {
		sc.clearBgColorFlags()
		sc.srgFlags.SetFlag(bgColor)
		sc.cacheValid = false
	}
	return sc
}

func (sc *SgrCsi) SetFgColor256(fgColor256 Color256) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlags(CustomFgColorFlag, FgColor256Flag)
	sc.fgColor256 = fgColor256
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetBgColor256(bgColor256 Color256) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlags(CustomBgColorFlag, BgColor256Flag)
	sc.bgColor256 = bgColor256
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetUnderlineColor256(underlineColor256 Color256) *SgrCsi {
	sc.clearUnderlineColorFlags()
	sc.srgFlags.SetFlags(CustomUnderlineColorFlag, UnderlineColor256Flag)
	sc.underlineColor256 = underlineColor256
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetFgColorRGB(fgColorRGB ColorRGB) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlags(CustomFgColorFlag, FgColorRGBFlag)
	sc.fgColorRGB = fgColorRGB
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetBgColorRGB(bgColorRGB ColorRGB) *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlags(CustomBgColorFlag, BgColorRGBFlag)
	sc.bgColorRGB = bgColorRGB
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetUnderlineColorRGB(underlineColorRGB ColorRGB) *SgrCsi {
	sc.clearUnderlineColorFlags()
	sc.srgFlags.SetFlags(CustomUnderlineColorFlag, UnderlineColorRGBFlag)
	sc.underlineColorRGB = underlineColorRGB
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetDefaultFgColor() *SgrCsi {
	sc.clearFgColorFlags()
	sc.srgFlags.SetFlag(DefaultFgColorFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetDefaultBgColor() *SgrCsi {
	sc.clearBgColorFlags()
	sc.srgFlags.SetFlag(DefaultBgColorFlag)
	sc.cacheValid = false
	return sc
}

func (sc *SgrCsi) SetDefaultUnderlineColor() *SgrCsi {
	sc.clearUnderlineColorFlags()
	sc.srgFlags.SetFlag(DefaultUnderlineColorFlag)
	sc.cacheValid = false
	return sc
}

func writeHeadAndSep(b *bytes.Buffer, headWritten *bool, needSep *bool) {
	if *needSep {
		b.WriteByte(';')
	} else if !*headWritten {
		b.WriteString(sgrCsiHead)
		*headWritten = true
	}
}

func writeSgr(b *bytes.Buffer, headWritten *bool, needSep *bool, sgr string) {
	writeHeadAndSep(b, headWritten, needSep)
	b.WriteString(sgr)
	*needSep = true
}

func writeSgrMulti(b *bytes.Buffer, headWritten *bool, needSep *bool, sgrMulti ...string) {
	if len(sgrMulti) == 0 {
		return
	}
	writeHeadAndSep(b, headWritten, needSep)
	b.WriteString(sgrMulti[0])
	for _, sgr := range sgrMulti[1:] {
		b.WriteByte(';')
		b.WriteString(sgr)
	}
}

func (sc *SgrCsi) writeFormat(b *bytes.Buffer, headWritten *bool, needSep *bool) {
	srgFlags := sc.srgFlags
	if s := formatTable1[srgFlags[1]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
	if s := formatTable2[srgFlags[2]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
	if s := formatTable3[srgFlags[3]&0b00111111]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
	if s := formatTable4[srgFlags[4]&0b00111111]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
	if s := formatTable5[srgFlags[5]&0b00111111]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
}

func (sc *SgrCsi) writeFgColor(b *bytes.Buffer, headWritten *bool, needSep *bool) {
	srgFlags := sc.srgFlags
	if srgFlags.TestFlag(DefaultFgColorFlag) {
		writeSgr(b, headWritten, needSep, defaultFgColorString)
		return
	}
	if srgFlags.TestFlag(CustomFgColorFlag) {
		switch srgFlags[11] & 0b00000011 {
		case 0b01:
			writeSgrMulti(b, headWritten, needSep, customFgColorString, "5", uint8ToStringTable[sc.fgColor256])
		case 0b10:
			writeSgrMulti(b, headWritten, needSep, customFgColorString, "2",
				uint8ToStringTable[sc.fgColorRGB.R],
				uint8ToStringTable[sc.fgColorRGB.G],
				uint8ToStringTable[sc.fgColorRGB.B],
			)
		}
		return
	}
	if s := fgNormalColorTable[sc.srgFlags[6]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
		return
	}
	if s := fgBrightColorTable[sc.srgFlags[8]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
}

func (sc *SgrCsi) writeBgColor(b *bytes.Buffer, headWritten *bool, needSep *bool) {
	srgFlags := sc.srgFlags
	if srgFlags.TestFlag(DefaultBgColorFlag) {
		writeSgr(b, headWritten, needSep, defaultBgColorString)
		return
	}
	if srgFlags.TestFlag(CustomBgColorFlag) {
		switch srgFlags[11] & 0b00000011 {
		case 0b01:
			writeSgrMulti(b, headWritten, needSep, customBgColorString, "5", uint8ToStringTable[sc.bgColor256])
		case 0b10:
			writeSgrMulti(b, headWritten, needSep, customBgColorString, "2",
				uint8ToStringTable[sc.bgColorRGB.R],
				uint8ToStringTable[sc.bgColorRGB.G],
				uint8ToStringTable[sc.bgColorRGB.B],
			)
		}
		return
	}
	if s := bgNormalColorTable[sc.srgFlags[7]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
		return
	}
	if s := bgBrightColorTable[sc.srgFlags[9]]; s != "" {
		writeSgr(b, headWritten, needSep, s)
	}
}

func (sc *SgrCsi) writeUnderlineColor(b *bytes.Buffer, headWritten *bool, needSep *bool) {
	srgFlags := sc.srgFlags
	if srgFlags.TestFlag(DefaultUnderlineColorFlag) {
		writeSgr(b, headWritten, needSep, defaultUnderlineColorString)
		return
	}
	if srgFlags.TestFlag(CustomUnderlineColorFlag) {
		switch srgFlags[11] & 0b00000011 {
		case 0b01:
			writeSgrMulti(b, headWritten, needSep, customUnderlineColorString, "5", uint8ToStringTable[sc.underlineColor256])
		case 0b10:
			writeSgrMulti(b, headWritten, needSep, customUnderlineColorString, "2",
				uint8ToStringTable[sc.underlineColorRGB.R],
				uint8ToStringTable[sc.underlineColorRGB.G],
				uint8ToStringTable[sc.underlineColorRGB.B],
			)
		}
	}
}

func (sc *SgrCsi) write(b *bytes.Buffer) {
	if sc.srgFlags.TestFlag(ResetFlag) {
		b.WriteString(resetSgrCsi)
		return
	}
	headWritten := false
	needSep := false
	sc.writeFormat(b, &headWritten, &needSep)
	sc.writeFgColor(b, &headWritten, &needSep)
	sc.writeBgColor(b, &headWritten, &needSep)
	sc.writeUnderlineColor(b, &headWritten, &needSep)
	if headWritten {
		b.WriteByte(sgrCsiEnd)
	}
}

func (sc *SgrCsi) RefreshCache() {
	buffer := &bytes.Buffer{}
	sc.write(buffer)
	sc.cache = uns.BytesToString(buffer.Bytes())
	sc.cacheValid = true
}

func (sc *SgrCsi) Write(b *bytes.Buffer) {
	if !sc.cacheValid {
		sc.RefreshCache()
	}
	b.WriteString(sc.cache)
}

func (sc *SgrCsi) WriteDisableCache(b *bytes.Buffer) {
	sc.write(b)
}

func (sc *SgrCsi) String() string {
	if !sc.cacheValid {
		sc.RefreshCache()
	}
	return sc.cache
}

func (sc *SgrCsi) StringDisableCache() string {
	b := &bytes.Buffer{}
	sc.write(b)
	return uns.BytesToString(b.Bytes())
}
