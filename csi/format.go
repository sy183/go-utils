package csi

import "strings"

var (
	//BoldFlag                 = 8
	//FaintFlag                = 9
	//ItalicFlag               = 10
	//FrakturFlag              = 11
	//ReversedFlag             = 12
	//NormalIntensityFlag      = 13
	//NotItalicBlackLetterFlag = 14
	//NotReversedFlag          = 15
	formatTable1 [256]string

	//UnderlineFlag        = 16
	//DoublyUnderlinedFlag = 17
	//SlowBlinkFlag        = 18
	//RapidBlinkFlag       = 19
	//OverlinedFlag        = 20
	//NotUnderlinedFlag    = 21
	//NotBlinkingFlag      = 22
	//NotOverlinedFlag     = 23
	formatTable2 [256]string

	//ConcealFlag                    = 24
	//CrossedOutFlag                 = 25
	//DisableProportionalSpacingFlag = 26
	//RevealFlag                     = 27
	//NotCrossedOutFlag              = 28
	//ProportionalSpacingFlag        = 29
	formatTable3 [256]string

	//FramedFlag                  = 32
	//EncircledFlag               = 33
	//SuperscriptFlag             = 34
	//SubscriptFlag               = 35
	//NotFramedEncircledFlag      = 36
	//NotSuperscriptSubscriptFlag = 37
	formatTable4 [256]string

	//IdeogramUnderlineFlag       = 40
	//IdeogramDoubleUnderlineFlag = 41
	//IdeogramOverlineFlag        = 42
	//IdeogramDoubleOverlineFlag  = 43
	//IdeogramStressMarkingFlag   = 44
	//NoIdeogramAttributesFlag    = 45
	formatTable5 [256]string
)

func initFormatTable1() {
	for i := 0; i < 256; i++ {
		var flags1S []string
		if i&0b00100000 != 0 {
			flags1S = append(flags1S, NormalIntensitySGR)
		} else if i&0b00000001 != 0 {
			flags1S = append(flags1S, BoldSGR)
		} else if i&0b00000010 != 0 {
			flags1S = append(flags1S, FaintSGR)
		}
		if i&0b01000000 != 0 {
			flags1S = append(flags1S, NotItalicBlackLetterSGR)
		} else if i&0b00000100 != 0 {
			flags1S = append(flags1S, ItalicSGR)
		} else if i&0b00001000 != 0 {
			flags1S = append(flags1S, FrakturSGR)
		}
		if i&0b10000000 != 0 {
			flags1S = append(flags1S, NotReversedSGR)
		} else if i&0b00010000 != 0 {
			flags1S = append(flags1S, ReversedSGR)
		}
		formatTable1[i] = strings.Join(flags1S, ";")
	}
}

func initFormatTable2() {
	for i := 0; i < 256; i++ {
		var flags2S []string
		if i&0b00100000 != 0 {
			flags2S = append(flags2S, NotUnderlinedSGR)
		} else if i&0b00000001 != 0 {
			flags2S = append(flags2S, UnderlineSGR)
		} else if i&0b00000010 != 0 {
			flags2S = append(flags2S, DoublyUnderlinedSGR)
		}
		if i&0b01000000 != 0 {
			flags2S = append(flags2S, NotBlinkingSGR)
		} else if i&0b00000100 != 0 {
			flags2S = append(flags2S, SlowBlinkSGR)
		} else if i&0b00001000 != 0 {
			flags2S = append(flags2S, RapidBlinkSGR)
		}
		if i&0b10000000 != 0 {
			flags2S = append(flags2S, NotOverlinedSGR)
		} else if i&0b00010000 != 0 {
			flags2S = append(flags2S, OverlinedSGR)
		}
		formatTable2[i] = strings.Join(flags2S, ";")
	}
}

func initFormatTable3() {
	for i := 0; i < 256; i++ {
		var flags3S []string
		if i&0b00001000 != 0 {
			flags3S = append(flags3S, RevealSGR)
		} else if i&0b00000001 != 0 {
			flags3S = append(flags3S, ConcealSGR)
		}
		if i&0b00010000 != 0 {
			flags3S = append(flags3S, NotCrossedOutSGR)
		} else if i&0b00000010 != 0 {
			flags3S = append(flags3S, CrossedOutSGR)
		}
		if i&0b00100000 != 0 {
			flags3S = append(flags3S, DisableProportionalSpacingSGR)
		} else if i&0b00000100 != 0 {
			flags3S = append(flags3S, ProportionalSpacingSGR)
		}
		formatTable3[i] = strings.Join(flags3S, ";")
	}
}

func initFormatTable4() {
	for i := 0; i < 256; i++ {
		var flags4S []string
		if i&0b00010000 != 0 {
			flags4S = append(flags4S, NotFramedEncircledSGR)
		} else if i&0b00000001 != 0 {
			flags4S = append(flags4S, FramedSGR)
		} else if i&0b00000010 != 0 {
			flags4S = append(flags4S, EncircledSGR)
		}
		if i&0b00100000 != 0 {
			flags4S = append(flags4S, NotSuperscriptSubscriptSGR)
		} else if i&0b00000100 != 0 {
			flags4S = append(flags4S, SuperscriptSGR)
		} else if i&0b00001000 != 0 {
			flags4S = append(flags4S, SubscriptSGR)
		}
		formatTable4[i] = strings.Join(flags4S, ";")
	}
}

func initFormatTable5() {
	for i := 0; i < 256; i++ {
		var flags5S []string
		if i&0b00100000 != 0 {
			flags5S = append(flags5S, NoIdeogramAttributesSGR)
		} else if i&0b00000001 != 0 {
			flags5S = append(flags5S, IdeogramUnderlineSGR)
		} else if i&0b00000010 != 0 {
			flags5S = append(flags5S, IdeogramDoubleUnderlineSGR)
		} else if i&0b00000100 != 0 {
			flags5S = append(flags5S, IdeogramOverlineSGR)
		} else if i&0b00001000 != 0 {
			flags5S = append(flags5S, IdeogramDoubleOverlineSGR)
		} else if i&0b00010000 != 0 {
			flags5S = append(flags5S, IdeogramStressMarkingSGR)
		}
		formatTable5[i] = strings.Join(flags5S, ";")
	}
}

func init() {
	initFormatTable1()
	initFormatTable2()
	initFormatTable3()
	initFormatTable4()
	initFormatTable5()
}
