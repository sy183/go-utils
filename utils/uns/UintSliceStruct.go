package uns

import "unsafe"

var uintSize = int(unsafe.Sizeof(uint(0)))

type UintSliceStruct SliceStruct

func UintSliceStructOf(us []uint) *UintSliceStruct {
	return (*UintSliceStruct)(unsafe.Pointer(&us))
}

func (uss *UintSliceStruct) UintSlice() []uint {
	return *(*[]uint)(unsafe.Pointer(uss))
}

func (uss *UintSliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: uss.Array, Len: uss.Len * uintSize, Cap: uss.Cap * uintSize}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (uss *UintSliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: uss.Array, Len: uss.Len * uintSize, Cap: uss.Cap * uintSize}
}

func UintSliceToBytes(us []uint) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&us))
	ss.Len *= uintSize
	ss.Cap *= uintSize
	return *(*[]byte)(unsafe.Pointer(ss))
}
