package uns

import "unsafe"

var uint32Size = int(unsafe.Sizeof(uint32(0)))

type Uint32SliceStruct SliceStruct

func Uint32SliceStructOf(us []uint32) *Uint32SliceStruct {
	return (*Uint32SliceStruct)(unsafe.Pointer(&us))
}

func (uss *Uint32SliceStruct) Uint32Slice() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(uss))
}

func (uss *Uint32SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: uss.Array, Len: uss.Len * uint32Size, Cap: uss.Cap * uint32Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (uss *Uint32SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: uss.Array, Len: uss.Len * uint32Size, Cap: uss.Cap * uint32Size}
}

func Uint32SliceToBytes(us []uint32) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&us))
	ss.Len *= uint32Size
	ss.Cap *= uint32Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
