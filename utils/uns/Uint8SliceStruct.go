package uns

import "unsafe"

var uint8Size = int(unsafe.Sizeof(uint8(0)))

type Uint8SliceStruct SliceStruct

func Uint8SliceStructOf(us []uint8) *Uint8SliceStruct {
	return (*Uint8SliceStruct)(unsafe.Pointer(&us))
}

func (uss *Uint8SliceStruct) Uint8Slice() []uint8 {
	return *(*[]uint8)(unsafe.Pointer(uss))
}

func (uss *Uint8SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: uss.Array, Len: uss.Len * uint8Size, Cap: uss.Cap * uint8Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (uss *Uint8SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: uss.Array, Len: uss.Len * uint8Size, Cap: uss.Cap * uint8Size}
}

func Uint8SliceToBytes(us []uint8) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&us))
	ss.Len *= uint8Size
	ss.Cap *= uint8Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
