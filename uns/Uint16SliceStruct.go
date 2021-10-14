package uns

import "unsafe"

var uint16Size = int(unsafe.Sizeof(uint16(0)))

type Uint16SliceStruct SliceStruct

func Uint16SliceStructOf(us []uint16) *Uint16SliceStruct {
	return (*Uint16SliceStruct)(unsafe.Pointer(&us))
}

func (uss *Uint16SliceStruct) Uint16Slice() []uint16 {
	return *(*[]uint16)(unsafe.Pointer(uss))
}

func (uss *Uint16SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: uss.Array, Len: uss.Len * uint16Size, Cap: uss.Cap * uint16Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (uss *Uint16SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: uss.Array, Len: uss.Len * uint16Size, Cap: uss.Cap * uint16Size}
}

func Uint16SliceToBytes(us []uint16) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&us))
	ss.Len *= uint16Size
	ss.Cap *= uint16Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
