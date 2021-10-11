package uns

import "unsafe"

var uint64Size = int(unsafe.Sizeof(uint64(0)))

type Uint64SliceStruct SliceStruct

func Uint64SliceStructOf(us []uint64) *Uint64SliceStruct {
	return (*Uint64SliceStruct)(unsafe.Pointer(&us))
}

func (uss *Uint64SliceStruct) Uint64Slice() []uint64 {
	return *(*[]uint64)(unsafe.Pointer(uss))
}

func (uss *Uint64SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: uss.Array, Len: uss.Len * uint64Size, Cap: uss.Cap * uint64Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (uss *Uint64SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: uss.Array, Len: uss.Len * uint64Size, Cap: uss.Cap * uint64Size}
}

func Uint64SliceToBytes(us []uint64) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&us))
	ss.Len *= uint64Size
	ss.Cap *= uint64Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
