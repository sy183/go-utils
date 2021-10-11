package flag

import (
	"fmt"
	"strings"
)

type BitFlags []byte

type BitFlag uint16

func (b BitFlag) Flag() string {
	panic("implement me")
}

func NewBitFlag(l uint16) BitFlags {
	return make(BitFlags, (int(l) - 1) / 8 + 1)
}

func (bf BitFlags) String() string {
	sbf := make([]string, len(bf))
	for i := 0; i < len(bf); i++ {
		sbf[i] = fmt.Sprintf("%08b", bf[i])
	}
	return "[" + strings.Join(sbf, " ") + "]"
}

func (bf BitFlags) SetFlag(flag BitFlag) {
	bf[flag>>3] |= 1 << (flag & 7)
}

func (bf BitFlags) UnsetFlag(flag BitFlag) {
	bf[flag>>3] &= ^(1 << (flag & 7))
}

func (bf BitFlags) SetFlags(flags ...BitFlag) {
	for _, flag := range flags {
		bf[flag>>3] |= 1 << (flag & 7)
	}
}

func (bf BitFlags) UnsetFlags(flags ...BitFlag) {
	for _, flag := range flags {
		bf[flag>>3] &= ^(1 << (flag & 7))
	}
}

func (bf BitFlags) Clean() {
	for i := 0; i < len(bf); i++ {
		bf[i] = 0
	}
}

func (bf BitFlags) TestFlag(flag BitFlag) bool {
	return (bf[flag>>3] & (1 << (flag & 7))) != 0
}

func (bf BitFlags) TestFlagUnset(flag BitFlag) bool {
	return (bf[flag>>3] & (1 << (flag & 7))) == 0
}

func (bf BitFlags) TestAllFlags(flags ...BitFlag) bool {
	for _, flag := range flags {
		if (bf[flag>>3] & (1 << (flag & 7))) == 0 {
			return false
		}
	}
	return true
}

func (bf BitFlags) TestAllFlagsUnset(flags ...BitFlag) bool {
	for _, flag := range flags {
		if (bf[flag>>3] & (1 << (flag & 7))) != 0 {
			return false
		}
	}
	return true
}

func (bf BitFlags) TestHasFlags(flags ...BitFlag) bool {
	for _, flag := range flags {
		if (bf[flag>>3] & (1 << (flag & 7))) != 0 {
			return true
		}
	}
	return false
}

func (bf BitFlags) TestHasFlagsUnset(flags ...BitFlag) bool {
	for _, flag := range flags {
		if (bf[flag>>3] & (1 << (flag & 7))) == 0 {
			return true
		}
	}
	return false
}
