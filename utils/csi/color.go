package csi

import (
	"strconv"
	"strings"
)

const (
	customFgColorString  = "38"
	defaultFgColorString = "39"
	customBgColorString  = "48"
	defaultBgColorString = "49"
)

type ColorRGB struct {
	R uint8
	G uint8
	B uint8
}

type Color256 uint8

type Color struct {
	normalFlags uint8
	brightFlags uint8
	customFlag  bool
	defaultFlag bool
	rgb256Flags uint8
	color256    Color256
	colorRGB    ColorRGB
}

var (
	fgNormalColorTable [256]string
	fgBrightColorTable [256]string
	bgNormalColorTable [256]string
	bgBrightColorTable [256]string
	uint8ToStringTable [256]string
)

func init() {
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 30)
		for j := 1 << i; j < 1<<(i+1); j++ {
			fgNormalColorTable[j] = s
		}
	}
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 40)
		for j := 1 << i; j < 1<<(i+1); j++ {
			bgNormalColorTable[j] = s
		}
	}
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 90)
		for j := 1 << i; j < 1<<(i+1); j++ {
			fgBrightColorTable[j] = s
		}
	}
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 100)
		for j := 1 << i; j < 1<<(i+1); j++ {
			bgBrightColorTable[j] = s
		}
	}
	for i := 0; i < 256; i++ {
		uint8ToStringTable[i] = strconv.Itoa(i)
	}
}

func (c *Color) fgString() string {
	if c.defaultFlag {
		return defaultFgColorString
	}
	if c.customFlag {
		switch c.rgb256Flags {
		case 0b01:
			return strings.Join([]string{customFgColorString, "5", uint8ToStringTable[c.color256]}, ";")
		case 0b10:
			return strings.Join([]string{customFgColorString, "2",
				uint8ToStringTable[c.colorRGB.R],
				uint8ToStringTable[c.colorRGB.G],
				uint8ToStringTable[c.colorRGB.B],
			}, ";")
		default:
			return ""
		}
	}
	s := fgNormalColorTable[c.normalFlags]
	if s != "" {
		return s
	}
	return fgBrightColorTable[c.brightFlags]
}

func (c *Color) bgString() string {
	if c.defaultFlag {
		return defaultBgColorString
	}
	if c.customFlag {
		switch c.rgb256Flags {
		case 0b01:
			return strings.Join([]string{customBgColorString, "5", uint8ToStringTable[c.color256]}, ";")
		case 0b10:
			return strings.Join([]string{customBgColorString, "2",
				uint8ToStringTable[c.colorRGB.R],
				uint8ToStringTable[c.colorRGB.G],
				uint8ToStringTable[c.colorRGB.B],
			}, ";")
		default:
			return ""
		}
	}
	s := bgNormalColorTable[c.normalFlags]
	if s != "" {
		return s
	}
	return bgBrightColorTable[c.brightFlags]
}
