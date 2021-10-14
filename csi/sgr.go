package csi

const (
	ResetFlag = 0

	BoldFlag                 = 8
	FaintFlag                = 9
	ItalicFlag               = 10
	FrakturFlag              = 11
	ReversedFlag             = 12
	NormalIntensityFlag      = 13
	NotItalicBlackLetterFlag = 14
	NotReversedFlag          = 15

	UnderlineFlag        = 16
	DoublyUnderlinedFlag = 17
	SlowBlinkFlag        = 18
	RapidBlinkFlag       = 19
	OverlinedFlag        = 20
	NotUnderlinedFlag    = 21
	NotBlinkingFlag      = 22
	NotOverlinedFlag     = 23

	ConcealFlag                    = 24
	CrossedOutFlag                 = 25
	DisableProportionalSpacingFlag = 26
	RevealFlag                     = 27
	NotCrossedOutFlag              = 28
	ProportionalSpacingFlag        = 29

	FramedFlag                  = 32
	EncircledFlag               = 33
	SuperscriptFlag             = 34
	SubscriptFlag               = 35
	NotFramedEncircledFlag      = 36
	NotSuperscriptSubscriptFlag = 37

	IdeogramUnderlineFlag       = 40
	IdeogramDoubleUnderlineFlag = 41
	IdeogramOverlineFlag        = 42
	IdeogramDoubleOverlineFlag  = 43
	IdeogramStressMarkingFlag   = 44
	NoIdeogramAttributesFlag    = 45

	FgBlackFlag   = 48
	FgRedFlag     = 49
	FgGreenFlag   = 50
	FgYellowFlag  = 51
	FgBlueFlag    = 52
	FgMagentaFlag = 53
	FgCyanFlag    = 54
	FgWhiteFlag   = 55

	BgBlackFlag   = 56
	BgRedFlag     = 57
	BgGreenFlag   = 58
	BgYellowFlag  = 59
	BgBlueFlag    = 60
	BgMagentaFlag = 61
	BgCyanFlag    = 62
	BgWhiteFlag   = 63

	FgBrightBlackFlag   = 64
	FgBrightRedFlag     = 65
	FgBrightGreenFlag   = 66
	FgBrightYellowFlag  = 67
	FgBrightBlueFlag    = 68
	FgBrightMagentaFlag = 69
	FgBrightCyanFlag    = 70
	FgBrightWhiteFlag   = 71

	BgBrightBlackFlag   = 72
	BgBrightRedFlag     = 73
	BgBrightGreenFlag   = 74
	BgBrightYellowFlag  = 75
	BgBrightBlueFlag    = 76
	BgBrightMagentaFlag = 77
	BgBrightCyanFlag    = 78
	BgBrightWhiteFlag   = 79

	CustomFgColorFlag         = 80
	DefaultFgColorFlag        = 81
	CustomBgColorFlag         = 82
	DefaultBgColorFlag        = 83
	CustomUnderlineColorFlag  = 84
	DefaultUnderlineColorFlag = 85

	FgColor256Flag        = 88
	FgColorRGBFlag        = 89
	BgColor256Flag        = 90
	BgColorRGBFlag        = 91
	UnderlineColor256Flag = 92
	UnderlineColorRGBFlag = 93

	PrimaryFontFlag      = 96
	AlternativeFont1Flag = 97
	AlternativeFont2Flag = 98
	AlternativeFont3Flag = 99
	AlternativeFont4Flag = 100
	AlternativeFont5Flag = 101
	AlternativeFont6Flag = 102
	AlternativeFont7Flag = 103

	AlternativeFont8Flag = 104
	AlternativeFont9Flag = 105
)

const (
	ResetSGR      = "0"
	BoldSGR       = "1"
	FaintSGR      = "2"
	ItalicSGR     = "3"
	UnderlineSGR  = "4"
	SlowBlinkSGR  = "5"
	RapidBlinkSGR = "6"
	ReversedSGR   = "7"
	ConcealSGR    = "8"
	CrossedOutSGR = "9"

	PrimaryFontSGR      = "10"
	AlternativeFont1SGR = "11"
	AlternativeFont2SGR = "12"
	AlternativeFont3SGR = "13"
	AlternativeFont4SGR = "14"
	AlternativeFont5SGR = "15"
	AlternativeFont6SGR = "16"
	AlternativeFont7SGR = "17"
	AlternativeFont8SGR = "18"
	AlternativeFont9SGR = "19"

	FrakturSGR              = "20"
	DoublyUnderlinedSGR     = "21"
	NormalIntensitySGR      = "22"
	NotItalicBlackLetterSGR = "23"
	NotUnderlinedSGR        = "24"
	NotBlinkingSGR          = "25"
	ProportionalSpacingSGR  = "26"
	NotReversedSGR          = "27"
	RevealSGR               = "28"
	NotCrossedOutSGR        = "29"

	FgBlackSGR        = "30"
	FgRedSGR          = "31"
	FgGreenSGR        = "32"
	FgYellowSGR       = "33"
	FgBlueSGR         = "34"
	FgMagentaSGR      = "35"
	FgCyanSGR         = "36"
	FgWhiteSGR        = "37"
	CustomFgColorSGR  = "38"
	DefaultFgColorSGR = "39"

	BgBlackSGR        = "40"
	BgRedSGR          = "41"
	BgGreenSGR        = "42"
	BgYellowSGR       = "43"
	BgBlueSGR         = "44"
	BgMagentaSGR      = "45"
	BgCyanSGR         = "46"
	BgWhiteSGR        = "47"
	CustomBgColorSGR  = "48"
	DefaultBgColorSGR = "49"

	DisableProportionalSpacingSGR = "50"
	FramedSGR                     = "51"
	EncircledSGR                  = "52"
	OverlinedSGR                  = "53"
	NotFramedEncircledSGR         = "54"
	NotOverlinedSGR               = "55"
	SetUnderlineColorSGR          = "58"
	DefaultUnderlineColorSGR      = "59"

	IdeogramUnderlineSGR       = "60"
	IdeogramDoubleUnderlineSGR = "61"
	IdeogramOverlineSGR        = "62"
	IdeogramDoubleOverlineSGR  = "63"
	IdeogramStressMarkingSGR   = "64"
	NoIdeogramAttributesSGR    = "65"

	SuperscriptSGR             = "73"
	SubscriptSGR               = "74"
	NotSuperscriptSubscriptSGR = "75"

	FgBrightBlackSGR   = "90"
	FgBrightRedSGR     = "91"
	FgBrightGreenSGR   = "92"
	FgBrightYellowSGR  = "93"
	FgBrightBlueSGR    = "94"
	FgBrightMagentaSGR = "95"
	FgBrightCyanSGR    = "96"
	FgBrightWhiteSGR   = "97"

	BgBrightBlackSGR   = "100"
	BgBrightRedSGR     = "101"
	BgBrightGreenSGR   = "102"
	BgBrightYellowSGR  = "103"
	BgBrightBlueSGR    = "104"
	BgBrightMagentaSGR = "105"
	BgBrightCyanSGR    = "106"
	BgBrightWhiteSGR   = "107"

	FgColor256SGR = "112"
	FgColorRGBSGR = "113"
	BgColor256SGR = "114"
	BgColorRGBSGR = "115"
)
