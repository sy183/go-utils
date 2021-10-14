package csi

import (
	"strconv"
)

const (
	customFgColorString         = "38"
	defaultFgColorString        = "39"
	customBgColorString         = "48"
	defaultBgColorString        = "49"
	customUnderlineColorString  = "48"
	defaultUnderlineColorString = "59"
)

type ColorRGB struct {
	R uint8
	G uint8
	B uint8
}

type Color256 uint8

var (
	fgNormalColorTable [256]string
	fgBrightColorTable [256]string
	bgNormalColorTable [256]string
	bgBrightColorTable [256]string
)

func initFgNormalColorTable() {
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 30)
		for j := 1 << i; j < 1<<(i+1); j++ {
			fgNormalColorTable[j] = s
		}
	}
}

func initFgBrightColorTable() {
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 90)
		for j := 1 << i; j < 1<<(i+1); j++ {
			fgBrightColorTable[j] = s
		}
	}
}

func initBgNormalColorTable() {
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 40)
		for j := 1 << i; j < 1<<(i+1); j++ {
			bgNormalColorTable[j] = s
		}
	}
}

func initBgBrightColorTable() {
	for i := 0; i < 8; i++ {
		s := strconv.Itoa(i + 100)
		for j := 1 << i; j < 1<<(i+1); j++ {
			bgBrightColorTable[j] = s
		}
	}
}

func init() {
	initFgNormalColorTable()
	initFgBrightColorTable()
	initBgNormalColorTable()
	initBgBrightColorTable()
}
